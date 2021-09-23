package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for i, value := range pow {
		fmt.Printf("%d == %d\n", i, value)
	}
}
