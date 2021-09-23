package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{"1", "2", "3", "4"}
	fmt.Println(names)

	a1 := names[0:2]
	b := names[1:3]
	fmt.Println(a1, b)

	b[0] = "X"
	fmt.Println(a1, b)
	fmt.Println(names)
}
