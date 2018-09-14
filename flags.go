package interrupts

type Flags struct {
	// Status flags
	S  bool // sign flag
	Z  bool // zero flag
	AC bool // auxiliary carry flag
	P  bool // parity flag
	CY bool // carry flag
	O  bool // overflow flag

	// Control flags
	D bool // directional flag
	I bool // interrupt flag
	T bool // trap flag
}
