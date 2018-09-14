package interrupts

type (
	Memory      [640 * 1024]uint8
	PixelBuffer []uint8 // 320x200
	Stack       []uint16
)

type State struct {
	reg *Registers
	mem *Memory
	s   *Stack
	f   *Flags
}
