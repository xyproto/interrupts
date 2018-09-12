package interrupts

import (
	"fmt"
)

type Registers struct {
	a, b, c, d, es, cs, di, ds *Register
}

// Fetch the values from the RegisterMap and update the registers
func merge(r *Registers, rm RegisterMap) {
	for k, v := range rm {
		switch k {
		case "ax":
			r.a.Set(v)
		case "ah":
			r.a.SetH(uint8(v))
		case "al":
			r.a.SetL(uint8(v))
		case "bx":
			r.b.Set(v)
		case "bh":
			r.b.SetH(uint8(v))
		case "bl":
			r.b.SetL(uint8(v))
		case "cx":
			r.c.Set(v)
		case "ch":
			r.c.SetH(uint8(v))
		case "cl":
			r.c.SetL(uint8(v))
		case "dx":
			r.d.Set(v)
		case "dh":
			r.d.SetH(uint8(v))
		case "dl":
			r.d.SetL(uint8(v))
		case "es":
			r.es.Set(v)
		case "cs":
			r.cs.Set(v)
		case "di":
			r.di.Set(v)
		case "ds":
			r.ds.Set(v)
		}
	}
}

// The function used for calling the reimplementations of interrupts
func Interrupt(n byte, r *Registers, m *Memory, s *Stack) {
	switch n {
	case 0x10:
		switch r.a.H() {
		case 0:
			rm := SetVideoMode(r.a.L())
			// TODO: read out values from rm and
			//       store in the global registers
			merge(r, rm)
		default:
			fmt.Printf("NOT IMPLEMENTED: INTERRUPT: %#x ah %#x\n", n, r.a.H())
		}
	case 0x13:
		fmt.Printf("CALLING INTERRUPT 0x13\n")
	case 0x19:
		fmt.Printf("CALLING INTERRUPT 0x19\n")
	case 0x21:
		fmt.Printf("CALLING INTERRUPT 0x21\n")
	default:
		fmt.Printf("NOT IMPLEMENTED: INTERRUPT: %#x\n", n)
	}
}
