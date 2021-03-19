package  main

import (
	"fmt"
)
//declare
var  y =  5
//declare empty int variable to indentifier "z"
// for 0 for booleans, 0 for integers , 0.0 for floats, "" for strings variable,
// and nil for pointers, functions, interfaces,slices,channels,  and maps.
var z int 
func  main() {

	x :=  10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	foo()

}
 func  foo() {
	 fmt.Println("again", y)
 }
