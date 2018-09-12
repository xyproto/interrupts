package interrupts

import (
	"encoding/binary"
)

type Register struct {
	l byte
	h byte
}

type RegisterMap map[string]uint16

func (r *Register) SetL(l uint8) {
	r.l = l
}

func (r *Register) SetH(h uint8) {
	r.h = h
}

func (r *Register) Set(w uint16) {
	bs := make([]uint8, 2)
	binary.LittleEndian.PutUint16(bs, w)
	r.l, r.h = bs[0], bs[1]
}

func (r *Register) SetR(e *Register) {
	r.l = e.l
	r.h = e.h
}

func (r *Register) Get() uint16 {
	return uint16(r.l) + uint16(r.h<<1)
}

func (r *Register) H() uint8 {
	return r.h
}

func (r *Register) L() uint8 {
	return r.l
}
