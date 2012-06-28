// Write a function which accepts 2 integers and returns thier sum, product and
// difference. Make a version with unnamed return variables, and a 2nd version
// with named return variables.
package main

import "fmt"

func main() {
	fmt.Println(names(2, 4))
	fmt.Println(noNames(2, 4))
}

func names(x, y int) (sum, product, difference int) {
	return (x + y), (x * y), (x - y)
}

func noNames(x, y int) (int, int, int) {
	return (x + y), (x * y), (x - y)
}
