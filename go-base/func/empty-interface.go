package main

import "fmt"

func describe1(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

func main() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)
}
