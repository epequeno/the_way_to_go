// Print the numbers from 10 to 1 in that order using a recursive function
// printrec(i int)
package main

import "fmt"

func main() {
	printrec(10)
}

func printrec(i int) {
	fmt.Println(i)
	if i > 1 {
		printrec(i - 1)
	}
}
