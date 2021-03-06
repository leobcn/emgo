// +build f411xe

package pwr

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type PWR_Periph struct {
	CR  CR
	CSR CSR
}

func (p *PWR_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var PWR = (*PWR_Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) LPDS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(LPDS)}}
}

func (p *PWR_Periph) PDDS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PDDS)}}
}

func (p *PWR_Periph) CWUF() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CWUF)}}
}

func (p *PWR_Periph) CSBF() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CSBF)}}
}

func (p *PWR_Periph) PVDE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PVDE)}}
}

func (p *PWR_Periph) PLS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLS)}}
}

func (p *PWR_Periph) DBP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBP)}}
}

func (p *PWR_Periph) FPDS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FPDS)}}
}

func (p *PWR_Periph) LPUDS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(LPUDS)}}
}

func (p *PWR_Periph) MRUDS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MRUDS)}}
}

func (p *PWR_Periph) ADCDC1() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADCDC1)}}
}

func (p *PWR_Periph) VOS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(VOS)}}
}

func (p *PWR_Periph) ODEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ODEN)}}
}

func (p *PWR_Periph) ODSWEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ODSWEN)}}
}

func (p *PWR_Periph) UDEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(UDEN)}}
}

func (p *PWR_Periph) FMSSR() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FMSSR)}}
}

func (p *PWR_Periph) FISSR() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FISSR)}}
}

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) WUF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WUF)}}
}

func (p *PWR_Periph) SBF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SBF)}}
}

func (p *PWR_Periph) PVDO() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PVDO)}}
}

func (p *PWR_Periph) BRR() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(BRR)}}
}

func (p *PWR_Periph) WUPP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WUPP)}}
}

func (p *PWR_Periph) EWUP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EWUP)}}
}

func (p *PWR_Periph) BRE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(BRE)}}
}

func (p *PWR_Periph) VOSRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VOSRDY)}}
}

func (p *PWR_Periph) ODRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ODRDY)}}
}

func (p *PWR_Periph) ODSWRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ODSWRDY)}}
}

func (p *PWR_Periph) UDSWRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(UDSWRDY)}}
}
