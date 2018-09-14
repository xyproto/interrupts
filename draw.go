package interrupts

import (
	"fmt"
	"time"
)

var (
	pixelBuffer PixelBuffer
	frameUpdate = 100 * time.Millisecond
)

// Set up graphics memory that will be used as a pixel buffer with SDL2
func Init(mem *Memory) {
	fmt.Println("CREATING A WINDOW")
	pixelBuffer = PixelBuffer(mem[0xa000 : 0xa000+(320*200)])
}

func Loop() {
	for {
		fmt.Println("UPDATING SCREEN")
		// TODO: Also read from keyboard
		// TODO: Support a palette as well
		//draw(dos.PixelBuffer(mem[0xa000:0xa000+(320*200)]))
		time.Sleep(frameUpdate)
	}
}

func Quit() {
	fmt.Println("CLOSING THE WINDOW")
}
