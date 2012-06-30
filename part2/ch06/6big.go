// Write a program which prints the factorial (!) of the first 30 integers
package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 10; i++ {
		j := int64(i)
		k := big.NewInt(j)
		fmt.Println(k)
	}
}

func fac(n *big.Int) (fac *big.Int) {
	testNum := big.NewInt(1)
	if n <= testNum {
		return testNum
	}
}
