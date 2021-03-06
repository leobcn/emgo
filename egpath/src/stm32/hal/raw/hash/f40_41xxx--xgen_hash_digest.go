// +build f40_41xxx

package hash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type HASH_DIGEST_Periph struct {
	HR [8]HR
}

func (p *HASH_DIGEST_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var HASH_DIGEST = (*HASH_DIGEST_Periph)(unsafe.Pointer(uintptr(mmap.HASH_DIGEST_BASE)))

type HR_Bits uint32

func (b HR_Bits) Field(mask HR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HR_Bits) J(v int) HR_Bits {
	return HR_Bits(bits.Make32(v, uint32(mask)))
}

type HR struct{ mmio.U32 }

func (r *HR) Bits(mask HR_Bits) HR_Bits { return HR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HR) StoreBits(mask, b HR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HR) SetBits(mask HR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HR) ClearBits(mask HR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HR) Load() HR_Bits             { return HR_Bits(r.U32.Load()) }
func (r *HR) Store(b HR_Bits)           { r.U32.Store(uint32(b)) }

type HR_Mask struct{ mmio.UM32 }

func (rm HR_Mask) Load() HR_Bits   { return HR_Bits(rm.UM32.Load()) }
func (rm HR_Mask) Store(b HR_Bits) { rm.UM32.Store(uint32(b)) }
