// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {
	// start run OMIT
	f := "apple"
	fmt.Println(f)
	f = "orange"
	fmt.Println(f)
	// end run OMIT
}
