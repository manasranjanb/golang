package main

import (
	"fmt"
)

type star int

var x star = 10
var y int = 20  

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 43
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}
