package main

import "fmt"

func main() {
	n,err := fmt.Println("hello world")
	fmt.Println(n)
	fmt.Println(err)
	foo()
}

func foo() {
      n,_ := fmt.Println("i'm foo")
      fmt.Println(n)

}
