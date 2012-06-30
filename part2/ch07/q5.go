/* Suppose we have the following slice 
items := [...]int{10, 20, 30, 40, 50}

a) if we code the following for-loop, what will be the value of items after 
the loop? 

for _, item := range items {
	item *= 2
}
Try it out if you are not sure.

Answer: the values double but the array is not modified.

b) If a) does not work, make a for loop in which the values are doubled. */
package main

import (
	"fmt"
)

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for ix, item := range items {
		items[ix] *= 2
		fmt.Println(item, items)
	}
}
