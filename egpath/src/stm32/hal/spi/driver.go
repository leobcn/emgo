package spi

import (
	"reflect"
	"rtos"
	"sync/atomic"
	"sync/fence"
	"unsafe"

	"stm32/hal/dma"
)

type DriverError byte

const ErrTimeout DriverError = 1

func (e DriverError) Error() string {
	switch e {
	case ErrTimeout:
		return "timeout"
	default:
		return ""
	}
}

type Driver struct {
	deadline int64

	P     *Periph
	RxDMA *dma.Channel
	TxDMA *dma.Channel

	dmacnt int
	done   rtos.EventFlag
	err    uint
}

// NewDriver provides convenient way to create heap allocated Driver struct.
func NewDriver(p *Periph, rxdma, txdma *dma.Channel) *Driver {
	d := new(Driver)
	d.P = p
	d.RxDMA = rxdma
	d.TxDMA = txdma
	return d
}

func (d *Driver) setupDMA(ch *dma.Channel, mode dma.Mode, wordSize uintptr) {
	ch.Setup(mode)
	ch.SetWordSize(wordSize, wordSize)
	ch.SetAddrP(unsafe.Pointer(d.P.raw.DR.U16.Addr()))
}

func (d *Driver) DMAISR(ch *dma.Channel) {
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	_, e := ch.Status()
	if e&^dma.ErrFIFO != 0 || atomic.AddInt(&d.dmacnt, -1) == 0 {
		d.done.Signal(1)
	}
}

func (d *Driver) ISR() {
	d.P.DisableIRQ(RxNotEmpty | Err)
	d.done.Signal(1)
}

func startDMA(ch *dma.Channel, addr uintptr, n int) {
	ch.SetAddrM(unsafe.Pointer(addr))
	ch.SetLen(n)
	ch.Clear(dma.EvAll, dma.ErrAll)
	ch.EnableIRQ(dma.Complete, dma.ErrAll&^dma.ErrFIFO) // Ignore FIFO error.
	fence.W()                                           // This orders writes to normal and I/O memory.
	ch.Enable()
}

func (d *Driver) SetDeadline(deadline int64) {
	d.deadline = deadline
}

func (d *Driver) WriteReadByte(b byte) byte {
	if d.err != 0 {
		return 0
	}
	d.P.SetDuplex(Full)
	d.P.EnableIRQ(RxNotEmpty | Err)
	d.done.Reset(0)
	fence.W() // This orders writes to normal and I/O memory.
	d.P.StoreByte(b)
	if !d.done.Wait(1, d.deadline) {
		d.err = uint(ErrTimeout) << 16
		return 0
	}
	b = d.P.LoadByte()
	if _, e := d.P.Status(); e != 0 {
		d.err = uint(e) << 8
		return 0
	}
	return b
}

func (d *Driver) writeReadDMA(out, in uintptr, olen, ilen int) int {
	txdmacfg := dma.MTP | dma.FIFO_4_4
	if olen > 1 {
		txdmacfg |= dma.IncM
	}
	d.setupDMA(d.TxDMA, txdmacfg, 1)
	d.setupDMA(d.RxDMA, dma.PTM|dma.IncM|dma.FIFO_1_4, 1)
	d.P.SetDuplex(Full)
	d.P.EnableDMA(RxNotEmpty | TxEmpty)
	d.P.EnableIRQ(Err)
	var n int
	for {
		m := ilen - n
		if m == 0 {
			return n
		}
		if m > 0xffff {
			m = 0xffff
		}
		d.dmacnt = 2
		d.done.Reset(0)
		startDMA(d.RxDMA, in, m)
		startDMA(d.TxDMA, out, m)
		if olen > 1 {
			out += uintptr(m)
		}
		in += uintptr(m)
		n += m

		done := d.done.Wait(1, d.deadline)

		d.P.DisableDMA(RxNotEmpty | TxEmpty)
		// Disable DMA (required by STM32F1 to allow setup next transfer).
		d.TxDMA.Disable()
		d.RxDMA.Disable()
		if !done {
			d.TxDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
			d.RxDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
			d.err = uint(ErrTimeout) << 16
			return n - d.RxDMA.Len()
		}
		if _, e := d.P.Status(); e != 0 {
			d.err = uint(e) << 8
			return n - d.RxDMA.Len()
		}
		_, rxe := d.RxDMA.Status()
		_, txe := d.TxDMA.Status()
		if e := (rxe | txe) &^ dma.ErrFIFO; e != 0 {
			d.err = uint(e)
			return n - d.RxDMA.Len()
		}
	}
}

func (d *Driver) writeDMA(out uintptr, olen int) {
	d.setupDMA(d.TxDMA, dma.MTP|dma.IncM|dma.FIFO_4_4, 1)
	d.P.SetDuplex(HalfOut) // Avoid ErrOverflow.
	d.P.EnableDMA(TxEmpty)
	d.P.EnableIRQ(Err)
	var n int
	for {
		m := olen - n
		if m == 0 {
			return
		}
		if m > 0xffff {
			m = 0xffff
		}
		d.dmacnt = 1
		d.done.Reset(0)
		startDMA(d.TxDMA, out+uintptr(n), m)
		n += m

		done := d.done.Wait(1, d.deadline)

		d.P.DisableDMA(TxEmpty)
		// Disable DMA (required by STM32F1 to allow setup next transfer).
		d.TxDMA.Disable()
		if !done {
			d.TxDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
			d.err = uint(ErrTimeout) << 16
			return
		}
		if _, e := d.P.Status(); e != 0 {
			d.err = uint(e) << 8
			return
		}
		_, txe := d.TxDMA.Status()
		if e := txe &^ dma.ErrFIFO; e != 0 {
			d.err = uint(e)
			return
		}
	}
}

// Err returns the first error that was encountered by the Driver. It also
// clears internal error flags so subsequent Err call returns nil or next error.
func (d *Driver) Err() error {
	e := d.err
	if e == 0 {
		return nil
	}
	d.err = 0
	if err := DriverError(e >> 16); err != 0 {
		return err
	}
	if err := Error(e >> 8); err != 0 {
		if err&ErrOverrun != 0 {
			d.P.LoadByte()
			d.P.Status()
		}
		return err
	}
	return dma.Error(e)
}

var ffff uint16 = 0xffff

func (d *Driver) WriteStringRead(out string, in []byte) int {
	olen := len(out)
	ilen := len(in)
	if d.err != 0 || olen == 0 && ilen == 0 {
		return 0
	}
	if olen <= 1 && ilen <= 1 {
		// Avoid DMA for one byte transfers.
		b := byte(0xff)
		if olen != 0 {
			b = out[0]
		}
		b = d.WriteReadByte(b)
		if ilen != 0 {
			in[0] = b
			return 1
		}
		return 0
	}
	oaddr := (*reflect.StringHeader)(unsafe.Pointer(&out)).Data
	iaddr := (*reflect.SliceHeader)(unsafe.Pointer(&in)).Data
	if olen > ilen {
		var n int
		if ilen > 0 {
			n = d.writeReadDMA(oaddr, iaddr, ilen, ilen)
			if d.err != 0 {
				return n
			}
			olen -= ilen
			oaddr += uintptr(ilen)
		}
		d.writeDMA(oaddr, olen)
		return n
	}
	if ilen > olen {
		var n int
		if olen > 0 {
			n = d.writeReadDMA(oaddr, iaddr, olen, olen)
			if d.err != 0 {
				return n
			}
			ilen -= olen
			iaddr += uintptr(olen)
			oaddr += uintptr(olen - 1)
		} else {
			oaddr = uintptr(unsafe.Pointer(&ffff))
		}
		return n + d.writeReadDMA(oaddr, iaddr, 1, ilen)
	}
	return d.writeReadDMA(oaddr, iaddr, ilen, ilen)
}

func (d *Driver) WriteRead(out, in []byte) int {
	return d.WriteStringRead(*(*string)(unsafe.Pointer(&out)), in)
}

func (d *Driver) WriteReadMany(oi ...[]byte) int {
	var n int
	for k := 0; k < len(oi); k += 2 {
		var in []byte
		if k+1 < len(oi) {
			in = oi[k+1]
		}
		out := oi[k]
		n += d.WriteRead(out, in)
	}
	return n
}
