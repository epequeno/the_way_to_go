// print out a rectangle of width = 20 and height = 10 with the * character
package main

import (
	"fmt"
)

func main() {
	makeRectangle(20, 10)
}

func makeRectangle(w, h int) {
	for i := 0; i < h; i++ {
		for i := 0; i < w; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
