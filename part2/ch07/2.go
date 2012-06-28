// Write the loop that fills an array with the loop-counter (from 0 to 15) and
// then prints that array to the screen
package main

import "fmt"

func main() {
	var a [16]int
	for i := 0; i < 16; i++ {
		a[i] = i
	}
	fmt.Println(a)
}
