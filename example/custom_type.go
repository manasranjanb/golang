package main

import (
	"fmt"
)

var y = 20

type star int

var u star = 10

func main() {
	fmt.Println(y)
	fmt.Println(u)
	fmt.Printf("%T\n", u)

}
