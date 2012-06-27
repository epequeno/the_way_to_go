/* Rewrite string_conversion2.go by using also := for the err variable, what
can changed? */

package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	var orig string = "123"
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig: %s is not an integer. Err: (%v)\n", orig, err)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n", an)
}
