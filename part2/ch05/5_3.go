// This program does not compile, why not?
// package main
// import "fmt"
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v\n", i)
// 	}
// 	fmt.Printf("%v\n", i) //<-- compile error: undefined i
// }
// It won't compile because we try to reference i outside of the loop body where
// it doesn't exit
//
// How could you make it work? 
// move the init statement outside the for loop so the entire func can see it.

package main

import "fmt"

func main() {
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("%v\n", i)
}
