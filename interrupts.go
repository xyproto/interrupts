package interrupts

import (
	"fmt"
)

// http://www.ctyme.com/intr/int.htm

// The function used for calling the reimplementations of interrupts
func Interrupt(n byte, state *State) {
	ah := state.reg.a.H()
	al := state.reg.a.L()
	switch n {
	case 0x10:
		switch ah {
		case 0:
			rm := SetVideoMode(al)
			// TODO: read out values from rm and
			//       store in the global registers
			merge(state.reg, rm)
		default:
			fmt.Printf("NOT IMPLEMENTED: INTERRUPT: %#x ah %#x\n", n, ah)
		}
	case 0x13:
		fmt.Printf("CALLING INTERRUPT 0x13\n")
	case 0x19:
		fmt.Printf("CALLING INTERRUPT 0x19\n")
	case 0x20:
		fmt.Printf("CALLING INTERRUPT 0x20\n")
	case 0x21:
		fmt.Printf("CALLING INTERRUPT 0x21\n")
	case 0x29:
		fmt.Printf("CALLING INTERRUPT 0x29\n")
	default:
		fmt.Printf("NOT IMPLEMENTED: INTERRUPT: %#x\n", n)
	}
}
