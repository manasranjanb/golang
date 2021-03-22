package main

import (
	"fmt"
)

func main() {

	x := "hello, how are you ?"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y := []byte(x)
	fmt.Println(y)
}
