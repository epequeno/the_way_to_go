// Split a buffer buf into 2 slices: the header is the first n bytes, the tail
// is the rest; use 1 line of code
package main

import (
	"fmt"
	"bytes"
)

var buf bytes.Buffer

func main() {
	data := "steven"
	buf.WriteString(data)
	splitBuff(2, buf)
}

func splitBuff(n int, buf bytes.Buffer) {
	fmt.Println(buf.String()[:n], " ", buf.String()[n:])
}
