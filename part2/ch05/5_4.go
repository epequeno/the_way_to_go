/* 1. Create a simple loop with the for construct. Make it loop 15 times and 
print out the loop counter with the fmt package.
2. Rewrite this loop using goto. The keyword for may not be used now. */
package main

import "fmt"

// func main() {
// 	for i := 1; i < 16; i++ {
// 		fmt.Println(i)
// 	}
// }
func main() {
	i := 0
	if i < 1 {
		goto F
	}
F:
	myPrinter(1)
}

func myPrinter(i int) {
	fmt.Println(i)
	if i < 15 {
		i++
		myPrinter(i)
	}

}

/* The book doesn't explain the 'goto' statement at all and the docs on it are 
kind of sparse as well. To me this is just silly, goto a function call which 
does the looping by recursion? I'm almost positive that isn't the solution the 
author expected but until I see more examples of the goto statement in use this
is what I can do. */
