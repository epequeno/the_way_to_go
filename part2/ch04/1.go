// This function will print "G", then "O" then "G" again since m() does not
// change the value of a. By using the ":=" initialization we create a local 
// a which is set to "O" leaving the global a as it was. When n() is called
// the second time it reads the unchanged, global a which is still "G"
// global a, which is now changed to "O"
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
	a := "O"
	print(a)
}
