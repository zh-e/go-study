package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)


func Walk(t *tree.Tree, ch chan int) {
	Visit(t, ch)
	close(ch)
}

func Visit(t *tree.Tree, ch chan int)  {
	ch <- t.Value
	if t.Right != nil {
		Visit(t.Right, ch)
	}

	if t.Left != nil {
		Visit(t.Left, ch)
	}
}

func Same(t1, t2  *tree.Tree) bool {
	i := 0
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for n := range ch1 {
		i ^= n
	}

	for n := range ch2 {
		i ^= n
	}

	return i == 0
}

func main() {
	t1, t2 := tree.New(1), tree.New(2)
	fmt.Println(Same(t1, t2))
}