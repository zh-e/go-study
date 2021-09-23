package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b",b)

	c := b[:2]
	printSlice1("c",c)

	d := c[2:5]
	printSlice1("d",d)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}