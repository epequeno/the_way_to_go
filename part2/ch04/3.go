// This program first run main() and sets a globally to "G" it calls f1() which 
// sets a locally to "O" and prints that, f1() calls f2() and f2() reads the
// global a which was set to "G" by main()
package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
