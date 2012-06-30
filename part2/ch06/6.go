// Write a program which prints the factorial (!) of the first 30 integers
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 31; i++ {
		n := float64(i)
		fmt.Printf("%v! = %v\n", n, factorial(n))
	}

}
func factorial(n float64) (fac float64) {
	if n == 0 || n == 1 {
		return 1
	}
	fac = n * factorial(n-1)
	return
}
