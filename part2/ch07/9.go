/*Given a slice s []int and a factor of type int, enlarge s so that its new 
length is len(s)*factor */
package main

import (
	"fmt"
)

func main() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(enlarge(arr, 2))
}

func enlarge(s []int, factor int) []int {
	myArr := make([]int, len(s)*factor)
	copy(myArr, s)
	return myArr
}
