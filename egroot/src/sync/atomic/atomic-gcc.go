// +build !cortexm0

package atomic

import "unsafe"

//c:static inline __attribute__((always_inline))
func compareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)

//c:static inline __attribute__((always_inline))
func compareAndSwapInt(addr *int, old, new int) (swapped bool)

//c:static inline __attribute__((always_inline))
func compareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)

//c:static inline __attribute__((always_inline))
func compareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

//c:static inline __attribute__((always_inline))
func swapUint32(addr *uint32, new uint32) (old uint32)

//c:static inline __attribute__((always_inline))
func swapInt(addr *int, new int) (old int)

//c:static inline __attribute__((always_inline))
func swapUintptr(addr *uintptr, new uintptr) (old uintptr)

//c:static inline __attribute__((always_inline))
func swapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

//c:static inline __attribute__((always_inline))
func addUint32(addr *uint32, delta uint32) (new uint32)

//c:static inline __attribute__((always_inline))
func addInt(addr *int, delta int) (new int)

//c:static inline __attribute__((always_inline))
func addUintptr(addr *uintptr, delta uintptr) (new uintptr)

//c:static inline __attribute__((always_inline))
func orUint32(addr *uint32, mask uint32) (new uint32)

//c:static inline __attribute__((always_inline))
func orUintptr(addr *uintptr, mask uintptr) (new uintptr)

//c:static inline __attribute__((always_inline))
func andUint32(addr *uint32, mask uint32) (new uint32)

//c:static inline __attribute__((always_inline))
func andUintptr(addr *uintptr, mask uintptr) (new uintptr)

//c:static inline __attribute__((always_inline))
func xorUint32(addr *uint32, mask uint32) (new uint32)

//c:static inline __attribute__((always_inline))
func xorUintptr(addr *uintptr, mask uintptr) (new uintptr)
