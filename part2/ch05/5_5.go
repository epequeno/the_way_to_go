/* Create a program that prints the following (up to 25 characters):
G
GG
GGG
GGGG
GGGGG
GGGGGG
GGGGGGG
1. use 2 nested for loops
2. use only one for loop and string concatenation
*/
package main

import (
	"fmt"
)

// func main() {
// 	for i := 1; i < 26; i++ {
// 		for j := 0; j < i; j++ {
// 			fmt.Print("G")
// 		}
// 		fmt.Println("")
// 	}
// }
func main() {
	word := "G"
	for i := 0; i < 26; i++ {
		fmt.Println(word)
		word += "G"
	}
}
