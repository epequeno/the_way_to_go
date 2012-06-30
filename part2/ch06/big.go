// practice with the math/big pack
package main

import (
	"fmt"
	"math/big"
	//"math"
)

func main() {
	num := big.NewInt(100)
	factor := big.NewInt(2)
	for i := 0; i < 10; i++ {
		num = num.Exp(num, factor, 1)
	}
	fmt.Println(num)
}
