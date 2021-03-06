// Peripheral: IWDG_Periph  Independent WATCHDOG.
// Instances:
//  IWDG  mmap.IWDG_BASE
// Registers:
//  0x00 32  KR
//  0x04 32  PR
//  0x08 32  RLR
//  0x0C 32  SR
// Import:
//  stm32/o/f10x_hd/mmap
package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	KEY KR_Bits = 0xFFFF << 0 //+ Key value (write only, read 0000h).
)

const (
	KEYn = 0
)

const (
	PR   PR_Bits = 0x07 << 0 //+ PR[2:0] (Prescaler divider).
	PR_0 PR_Bits = 0x01 << 0 //  Bit 0.
	PR_1 PR_Bits = 0x02 << 0 //  Bit 1.
	PR_2 PR_Bits = 0x04 << 0 //  Bit 2.
)

const (
	PRn = 0
)

const (
	RL RLR_Bits = 0xFFF << 0 //+ Watchdog counter reload value.
)

const (
	RLn = 0
)

const (
	PVU SR_Bits = 0x01 << 0 //+ Watchdog prescaler value update.
	RVU SR_Bits = 0x01 << 1 //+ Watchdog counter reload value update.
)

const (
	PVUn = 0
	RVUn = 1
)
