// Rewrite the Fibonacci program above to return 2 named variables, namely the 
// value and the position of the Fibonacci-number, like 5 and 4 or 89 and 10
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 11; i++ {
		result, pos := fibonacci(i)
		fmt.Printf("pos %d, res %d\n", pos, result)
	}

}

func fibonacci(n int) (res, pos int) {
	if n < 1 {
		res = 0
	} else if n == 1 {
		res = 1
	} else {
		one, _ := fibonacci(n - 1)
		two, _ := fibonacci(n - 2)
		res = one + two
	}
	return res, n
}
