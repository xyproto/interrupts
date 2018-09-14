package interrupts

import (
	"encoding/binary"
)

type Register struct {
	l byte
	h byte
}

type RegisterMap map[string]uint16

type Registers struct {
	a, b, c, d, si, di, sp, bp, ip, cs, es, ds, fs, gs, ss *Register
}

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

func (r *Registers) A() *Register {
	return r.a
}

func (r *Registers) B() *Register {
	return r.b
}

func (r *Registers) C() *Register {
	return r.c
}

func (r *Registers) D() *Register {
	return r.d
}

func (r *Registers) SI() *Register {
	return r.si
}

func (r *Registers) DI() *Register {
	return r.di
}

func (r *Registers) SP() *Register {
	return r.sp
}

func (r *Registers) BP() *Register {
	return r.bp
}

func (r *Registers) IP() *Register {
	return r.ip
}

func (r *Registers) CS() *Register {
	return r.cs
}

func (r *Registers) ES() *Register {
	return r.es
}

func (r *Registers) DS() *Register {
	return r.ds
}

func (r *Registers) FS() *Register {
	return r.fs
}

func (r *Registers) GS() *Register {
	return r.gs
}

func (r *Registers) SS() *Register {
	return r.ss
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
		case "si":
			r.si.Set(v)
		case "di":
			r.di.Set(v)
		case "sp":
			r.sp.Set(v)
		case "bp":
			r.bp.Set(v)
		case "ip":
			r.ip.Set(v)
		case "cs":
			r.cs.Set(v)
		case "es":
			r.es.Set(v)
		case "ds":
			r.ds.Set(v)
		case "fs":
			r.fs.Set(v)
		case "gs":
			r.gs.Set(v)
		case "ss":
			r.ss.Set(v)
		}
	}
}
