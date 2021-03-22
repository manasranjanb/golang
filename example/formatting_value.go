package main

import (
	"fmt"
)

var y = 14

func main() {

	fmt.Println(y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#b\n", y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%o\n", y)
	fmt.Print(y)
}
