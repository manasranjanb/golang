package main

import (
	"fmt"
)

type star int

var x star = 10

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
