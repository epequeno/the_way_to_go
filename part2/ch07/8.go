/* Write a minSlice function which takes a slice of ints and returns the 
minimum, and a maxSlice function which takes a slice of ints and returns the 
maximum */
package main

import (
	"fmt"
)

func main() {
	arr := [...]int{5, 4, 1, 7, 4, 10, 2, 4, 1, 6, 0}
	fmt.Println(minSlice(arr[:]))
	fmt.Println(maxSlice(arr[:]))

}

func minSlice(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
		}
	}
	return min
}

func maxSlice(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
		}
	}
	return max
}
