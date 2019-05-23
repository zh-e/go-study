package main

import "fmt"

func main() {
	forDemo()
	forDemo1()
	forDemo2()
	ifDemo(9)
	ifDemo1(8)
	switchDemo()
	deferDemo()
	deferDemo1()
}

func forDemo() {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forDemo1() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

//go中用for替换while
func forDemo2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func ifDemo(x int) {
	if x > 0 {
		fmt.Println(x, "大于 0")
		return
	}
	fmt.Println(x, "小于等于 0")
}

func ifDemo1(x int) {
	//n 变量的作用域仅在if语句内
	if n := 9; x < n {
		fmt.Println(x, "小于", n)
	} else if x == n {
		fmt.Println(x, "等于", n)
	} else {
		fmt.Println(x, "大于", n)
	}
}

func switchDemo() {
	n := 4
	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
	case 5:
		fmt.Println(4, 5)
	default:
		fmt.Println(0)
	}
}

//defer 语句会将函数推迟到外层函数返回之后执行。
func deferDemo() {
	defer fmt.Println("111")
	fmt.Println("22222")
}

//后进先出顺序
func deferDemo1() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}