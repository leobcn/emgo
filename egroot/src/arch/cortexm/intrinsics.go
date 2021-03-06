// Package cortexm provides intrinsic functions for many Cortex-M special
// instructions.
package cortexm

import "unsafe"

// SEV inserts Signal Event instruction.
//
// Implies compiler fence.
//
//c:static inline
func SEV()

// DMB inserts Data Memory Barrier instruction.
//
// Implies compiler fence.
//
//c:static inline
func DMB()

// DSB inserts Data Synchronization Barrier instruction.
//
// Implies compiler fence.
//
//c:static inline
func DSB()

// ISB inserts Instruction Synchronization Barrier instruction.
//
// Implies compiler fence.
//
//c:static inline
func ISB()

// WFE inserts Wait For Event instruction.
//
// Implies compiler fence.
//
//c:static inline
func WFE()

// WFI inserts Wait For Interrupt instruction.
//
// Implies compiler fence.
//
//c:static inline
func WFI()

// SVC inserts Supervisor Call instruction.
//
// Implies compiler fence.
func SVC(imm byte)

// BKPT inserts Breakpoint instruction.
//
// Implies compiler fence.
func BKPT(imm byte)

// PRIMASK returns true if all exceptions with configurable priority are
// disabled.
//
//c:static inline
func PRIMASK() bool

// SetPRIMASK disables all exceptions with configurable priority. Internally it
// inserts cpsid i instruction. Atomic primitives on Cortex-M0 always enable
// exceptions after atomic operation. If you need this functions on Cortex-M0
// you should be very careful.
//
// Implies compiler fence.
//
//c:static inline
func SetPRIMASK()

// ClearPRIMASK reverts SetPRIMASK. Internally it inserts cpsie i instruction.
// If you modified any data that can be used by enabled interrupt handlers you
// probably need to call fence.Memory() before use this function.
//
// Implies compiler fence.
//
//c:static inline
func ClearPRIMASK()

// FAULTMASK returns true if all exceptions other than NMI are disabled.
//
//c:static inline
func FAULTMASK() bool

// SetFAULTMASK disables all exceptions other than NMI. Internally it inserts
// cpsid f instruction. Not supported by Cortex-M0.
//
// Implies compiler fence.
//
//c:static inline
func SetFAULTMASK()

// ClearFAULTMASK reverts SetFAULTMASK. Internally it inserts cpsie f
// instruction. If you modified any data that can be used by enabled interrupt
// handlers you probably need to call fence.Memory() before use this function.
// Not supported by Cortex-M0.
//
// Implies compiler fence.
//
//c:static inline
func ClearFAULTMASK()

// BASEPRIO returns current value of BASEPRI register.
//
//c:static inline
func BASEPRI() byte

// SetBASEPRI sets BASEPRI register. It prevents the activation of exceptions
// with the same or lower as p. Not supported by Cortex-M0.
//
// Implies compiler fence.
//
//c:static inline
func SetBASEPRI(p byte)

//c:static inline
func APSR() uint32

//c:static inline
func SetAPSR(r uint32)

//c:static inline
func IPSR() uint32

//c:static inline
func EPSR() uint32

//c:static inline
func IEPSR() uint32

//c:static inline
func IAPSR() uint32

//c:static inline
func EAPSR() uint32

//c:static inline
func PSR() uint32

//c:static inline
func SetPSR(r uint32)

//c:static inline
func MSP() uintptr

//c:static inline
func SetMSP(p unsafe.Pointer)

//c:static inline
func PSP() uintptr

//c:static inline
func SetPSP(p unsafe.Pointer)

//c:static inline
func LR() uint32

//c:static inline
func SetLR(r uint32)

type Cflags uint32

const (
	Unpriv Cflags = 1 << 0 // Unprivileged thread mode
	UsePSP Cflags = 1 << 1 // Use PSP for in thread mode
	FPCA   Cflags = 1 << 2 // Floating-point context active
)

//c:static inline
func SetCONTROL(c Cflags)

//c:static inline
func CONTROL() Cflags
