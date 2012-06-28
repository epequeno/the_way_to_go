// What will this loop print out
// for i := 0; i < 5; i++ {
// 	var v int
// 	fmt.Printf("%d ", v)
// 	v = 5
// }
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

/* Since the variable is initiated inside the for loop, it's value is set to
the default value for integers: 0. Had the variable been set outside the for
loop, the variable would be set to 5 after the first iteration and would then
print 5's */
