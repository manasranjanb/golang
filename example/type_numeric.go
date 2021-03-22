package main

import (
	"fmt"
)

var x = 42
var y = 43.002

func main() {
	x := 1
	y = 2

	fmt.Println(x, y)
	fmt.Printf("%T\n%T\n", x, y)
}
