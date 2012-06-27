/*
Write a program that prints the numbers from 1 to 100, but for multiples 
of three print "Fizz" instead of the number and for the multiples of 5 five
print "Buzz". For the numbers which are multiples of both three and five print
"FizzBuzz". (hint: use a switch with conditions)
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 101; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
