// Peripheral: DBGMCU_Periph  Debug MCU.
// Instances:
//  DBGMCU  mmap.DBGMCU_BASE
// Registers:
//  0x00 32  IDCODE MCU device ID code.
//  0x04 32  CR     Debug MCU configuration register.
//  0x08 32  APB1FZ Debug MCU APB1 freeze register.
//  0x0C 32  APB2FZ Debug MCU APB2 freeze register.
// Import:
//  stm32/o/f40_41xxx/mmap
package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	DEV_ID IDCODE_Bits = 0xFFF << 0   //+
	REV_ID IDCODE_Bits = 0xFFFF << 16 //+
)

const (
	DEV_IDn = 0
	REV_IDn = 16
)

const (
	DBG_SLEEP    CR_Bits = 0x01 << 0 //+
	DBG_STOP     CR_Bits = 0x01 << 1 //+
	DBG_STANDBY  CR_Bits = 0x01 << 2 //+
	TRACE_IOEN   CR_Bits = 0x01 << 5 //+
	TRACE_MODE   CR_Bits = 0x03 << 6 //+
	TRACE_MODE_0 CR_Bits = 0x01 << 6 //  Bit 0.
	TRACE_MODE_1 CR_Bits = 0x02 << 6 //  Bit 1.
)

const (
	DBG_SLEEPn   = 0
	DBG_STOPn    = 1
	DBG_STANDBYn = 2
	TRACE_IOENn  = 5
	TRACE_MODEn  = 6
)
