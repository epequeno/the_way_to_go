// Define an alias type Rope for string and declare a variable with it.
package main

import (
	"fmt"
)

type Rope string

func main() {
	var name Rope = "Steven"
	fmt.Println(name)
}
