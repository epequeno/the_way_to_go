/* Make a function that has as arguments a variable number of ints and
which prints each int on a seperatle line */
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	printInts(arr...)
}

func printInts(ints ...int) {
	if len(ints) == 0 {
		fmt.Println("No ints to print!")
	}
	for _, value := range ints {
		fmt.Println(value)
	}
}
