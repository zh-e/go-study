package main

import "fmt"

func main() {
	var s []int
	printSlice2(s)

	s =append(s, 0)
	printSlice2(s)

	s = append(s, 1)
	printSlice2(s)

	s = append(s, 2, 3, 5)
	printSlice2(s)

}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
