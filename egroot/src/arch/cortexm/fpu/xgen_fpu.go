package fpu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"
)


func fpu(n uint) *mmio.U32 {
	return &(*[3]mmio.U32)(unsafe.Pointer(uintptr(0xe000ef34)))[n]
}


type FPCCR_Bits uint32

func (m FPCCR_Bits) Load() uint32   { return fpu(0).Bits(uint32(m)) }
func (m FPCCR_Bits) Store(b uint32) { fpu(0).StoreBits(uint32(m), b) }
func (m FPCCR_Bits) Set()           { fpu(0).SetBits(uint32(m)) }
func (m FPCCR_Bits) Clear()         { fpu(0).ClearBits(uint32(m)) }

type FPCCR_Field uint16

func (f FPCCR_Field) Load() int   { return fpu(0).Field(uint16(f)) }
func (f FPCCR_Field) Store(v int) { fpu(0).SetField(uint16(f), v) }

func FPCCR_Load() FPCCR_Bits   { return FPCCR_Bits(fpu(0).Load()) }
func FPCCR_Store(b FPCCR_Bits) { fpu(0).Store(uint32(b)) }


type FPCAR_Bits uint32

func (m FPCAR_Bits) Load() uint32   { return fpu(1).Bits(uint32(m)) }
func (m FPCAR_Bits) Store(b uint32) { fpu(1).StoreBits(uint32(m), b) }
func (m FPCAR_Bits) Set()           { fpu(1).SetBits(uint32(m)) }
func (m FPCAR_Bits) Clear()         { fpu(1).ClearBits(uint32(m)) }

type FPCAR_Field uint16

func (f FPCAR_Field) Load() int   { return fpu(1).Field(uint16(f)) }
func (f FPCAR_Field) Store(v int) { fpu(1).SetField(uint16(f), v) }

func FPCAR_Load() FPCAR_Bits   { return FPCAR_Bits(fpu(1).Load()) }
func FPCAR_Store(b FPCAR_Bits) { fpu(1).Store(uint32(b)) }


type FPDSCR_Bits uint32

func (m FPDSCR_Bits) Load() uint32   { return fpu(2).Bits(uint32(m)) }
func (m FPDSCR_Bits) Store(b uint32) { fpu(2).StoreBits(uint32(m), b) }
func (m FPDSCR_Bits) Set()           { fpu(2).SetBits(uint32(m)) }
func (m FPDSCR_Bits) Clear()         { fpu(2).ClearBits(uint32(m)) }

type FPDSCR_Field uint16

func (f FPDSCR_Field) Load() int   { return fpu(2).Field(uint16(f)) }
func (f FPDSCR_Field) Store(v int) { fpu(2).SetField(uint16(f), v) }

func FPDSCR_Load() FPDSCR_Bits   { return FPDSCR_Bits(fpu(2).Load()) }
func FPDSCR_Store(b FPDSCR_Bits) { fpu(2).Store(uint32(b)) }
