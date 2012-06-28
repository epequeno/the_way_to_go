// Prove that when assigning an array to another, a distinct copy in memory of
// the array is made.
package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Printf("The array a is Type %T and the address is %p\n", a, &a)
	b := a
	fmt.Printf("The array b is Type %T and the address is %p\n", b, &b)
}
