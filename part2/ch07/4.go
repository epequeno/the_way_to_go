// Starting from solution Ex 7.3, make a version in which main calls
// a function with parameter the number of terms in the series. The function
// returns a slice with the Fibonacci numbers up to that number.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibo(15))
}

// The func fibo(n) takes an integer n as its single argument and returns 
// fibn a slice of unint64 whose len = (len(n)+3), three spaces of padding
// to compensate for the offset caused by populating the inital fibo array (arr)
// with the first 3 in the fibo sequence.
func fibo(n int) (fibn []uint64) {
	fibn = make([]uint64, n+3)
	arr := [3]uint64{0, 1, 1}
	for i := 0; i < 3; i++ {
		fibn[i] = arr[i]
	}
	for i := 0; i < n; i++ {
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[0] + arr[1]
		fibn[i+3] = arr[2]
	}
	return fibn[:len(fibn)-2]
}
