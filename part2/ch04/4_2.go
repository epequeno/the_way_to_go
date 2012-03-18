// This function will print "G", then "O" then "O" again since m() changes
// the value of a globally with the "=" assignment. When n() is called the
// second time it reads the global a, which is now changed to "O"
package main
var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
