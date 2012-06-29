// Given a slice sl we want to append a []byte data.
// Write a function Append(slice, data[]byte) []byte which lets sl grow if it
// is not big enough to accomodate data.
package main

import (
	"fmt"
)

var sl = make([]byte, 3)
var bytes = make([]byte, 10)

func main() {
	// Begin dummy data
	for i := 0; i < 10; i++ {
		bytes[i] = byte(i)
	}

	for i := 0; i < 3; i++ {
		sl[i] = byte(i + 100)
	}
	//end

	newsl := Append(sl, bytes)

	fmt.Printf("sl: %v len: %d\n", sl, len(sl))
	fmt.Printf("bytes: %v len: %d\n", bytes, len(bytes))
	fmt.Printf("new sl: %v len: %d\n", newsl, len(newsl))

}

func Append(slice, data []byte) []byte {
	mySlice := make([]byte, len(data)+len(slice))
	for i := 0; i < len(data); i++ {
		if i < len(slice) {
			mySlice[i] = slice[i]
		}
		mySlice[i+(len(slice))] = data[i]
	}
	return mySlice
}
