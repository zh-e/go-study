package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(50 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("      .")
			time.Sleep(5 * time.Millisecond)

		}
	}
}
