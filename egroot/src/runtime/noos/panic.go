// +build cortexm0 cortexm3 cortexm4 cortexm4f cortexm7f cortexm7d
package noos

import (
	"arch/cortexm"
	//"arch/cortexm/debug/itm"
)

/*
const dbg = itm.Port(0)

type stringer interface {
	String() string
}
*/

func panic_(i interface{}) {
	/*
		var s string
		switch v := i.(type) {
		case error:
			s = v.Error()
		case stringer:
			s = v.String()
		case string:
			s = v
		default:
		}
		dbg.WriteString("panic: ")
		dbg.WriteString(s)
		dbg.WriteByte('\n')
	*/
	for {
		cortexm.BKPT(1)
	}
}
