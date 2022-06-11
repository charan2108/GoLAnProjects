// package main tells  compiler that the
//  package should be compiled as an executable program.
// instead of main library
// main function in the package will be the entry point of executable program
// when you build shred libraries you will not have a main package and
// main function in the package
package main
// import fmt
// the keyword import is used to import a package to another packages
// Fmt stands for  format package which allows to format basic string values 
// or  anything and print them or collect user input from console
// write into a file using writer print customized fancy error
// used for formatting input and output
import "fmt"

func main() {
	fmt.Println("Welcome to Go Lang")
}
