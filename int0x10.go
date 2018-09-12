package interrupts

import (
	"fmt"
)

func SetVideoMode(al byte) RegisterMap {
	fmt.Println("SET VIDEO MODE", al)
	return RegisterMap{"al": 0}
}

func SetTextModeCursorShape(ch, cl byte) {
	fmt.Println("SET TEXT-MODE CURSOR SHAPE", cl, ch)
}

func SetCursorPosition(bh, dh, dl byte) {
	fmt.Println("SET CURSOR POSITION", bh, dl, dh)
}

func GetCursorPositionAndShape(bh byte) RegisterMap {
	fmt.Println("GET CURSOR POSITION AND SHAPE")
	return RegisterMap{"ax": 0, "ch": 0, "cl": 0, "dh": 0, "dl": 0}
}

func ReadLightPenPosition() RegisterMap {
	fmt.Println("READ LIGHT PEN POSITION")
	return RegisterMap{"ah": 0, "bx": 0, "cx": 0, "dh": 0, "dl": 0}
}

func SelectActiveDisplayPage(al byte) {
	fmt.Println("SELECT ACTIVE DISPLAY PAGE")
}

func ScrollUpWindow(al, bh, ch, cl, dh, dl byte) {
	fmt.Println("SCROLL UP WINDOW")
}

func ScrollDownWindow(al, bh, ch, cl, dh, dl byte) {
	fmt.Println("SCROLL DOWN WINDOW")
}

// More to implement:
// https://en.wikipedia.org/wiki/INT_10H
