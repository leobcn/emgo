// Package timer provides interface to manage nRF51 timers.
package timer

import (
	"mmio"
	"unsafe"

	"nrf51/hal/internal"
	"nrf51/hal/te"
)

// Periph represents timer/counter peripheral.
type Periph struct {
	te.Regs

	_         [65]mmio.U32
	mode      mmio.U32 // Timer mode selection.
	bitmode   mmio.U32 // Configure the number of bits used by the TIMER.
	_         mmio.U32
	prescaler mmio.U32 // Timer prescaler register.
	_         [11]mmio.U32
	cc        [4]mmio.U32 // Capture/Compare registers.
}

//emgo:const
var (
	Timer0 = (*Periph)(unsafe.Pointer(internal.BaseAPB + 0x08000))
	Timer1 = (*Periph)(unsafe.Pointer(internal.BaseAPB + 0x09000))
	Timer2 = (*Periph)(unsafe.Pointer(internal.BaseAPB + 0x0a000))
)

type TASK int

const (
	START    TASK = 0  // Start Timer.
	STOP     TASK = 1  // Stop Timer.
	COUNT    TASK = 2  // Increment Timer (Counter mode only).
	CLEAR    TASK = 3  // Clear timer.
	CAPTURE0 TASK = 16 // Capture Timer value to CC0 register.
	CAPTURE1 TASK = 17 // Capture Timer value to CC1 register.
	CAPTURE2 TASK = 18 // Capture Timer value to CC2 register.
	CAPTURE3 TASK = 19 // Capture Timer value to CC3 register.
)

type EVENT int

const (
	COMPARE0 EVENT = 16 // Compare event on CC[0] match.
	COMPARE1 EVENT = 17 // Compare event on CC[1] match.
	COMPARE2 EVENT = 18 // Compare event on CC[2] match.
	COMPARE3 EVENT = 19 // Compare event on CC[3] match.
)

type Shorts uint32

const (
	COMPARE0_CLEAR Shorts = 1 << 0
	COMPARE1_CLEAR Shorts = 1 << 1
	COMPARE2_CLEAR Shorts = 1 << 2
	COMPARE3_CLEAR Shorts = 1 << 3
	COMPARE0_STOP  Shorts = 0x100 << 0
	COMPARE1_STOP  Shorts = 0x100 << 1
	COMPARE2_STOP  Shorts = 0x100 << 2
	COMPARE3_STOP  Shorts = 0x100 << 3
)

func (p *Periph) Task(t TASK) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e EVENT) *te.Event { return p.Regs.Event(int(e)) }

//func (p *Periph) SHORTS() *ShortsReg { return (*ShortsReg)(unsafe.Pointer(p.ph.Shorts.Addr())) }

type Mode byte

const (
	TIMER   Mode = 0
	COUNTER Mode = 1
)

func (p *Periph) Mode() Mode {
	return Mode(p.mode.Bits(1))
}

func (p *Periph) SetMode(m Mode) {
	p.mode.Store(uint32(m))
}

type Bitmode byte

const (
	BIT8  Bitmode = 1
	BIT16 Bitmode = 0
	BIT24 Bitmode = 2
	BIT32 Bitmode = 3
)

func (p *Periph) Bitmode() Bitmode {
	return Bitmode(p.bitmode.Bits(3))
}

func (p *Periph) SetBitmode(m Bitmode) {
	p.bitmode.Store(uint32(m))
}

func (p *Periph) Prescaler() int {
	return int(p.prescaler.Bits(0xf))
}

// SetPrescaler sets prescaler to exp (freq = 16MHz/2^exp). Must only be used
// when the timer is stopped.
func (p *Periph) SetPrescaler(exp int) {
	p.prescaler.Store(uint32(exp))
}

func (p *Periph) CC(n int) uint32 {
	return p.cc[n].Load()
}

func (p *Periph) SetCC(n int, cc uint32) {
	p.cc[n].Store(cc)
}
