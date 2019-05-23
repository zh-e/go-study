package main

import "fmt"

func main() {
	i, j := 42, 2701

	//指向i
	p := &i
	// 通过指针读取i的值
	fmt.Println(*p)
	//通过指针更改i的值
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}


