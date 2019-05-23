package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	random()
	sqrt()
	fmt.Println(add1(1, 2))
	fmt.Println(swap("qqq", "zzzz"))
	fmt.Println(split(18))
}

//此方法会引入fmt math/rand
func random() {
	fmt.Println("num: ", rand.Intn(100))
}

func sqrt() {
	fmt.Println("sqrt: ", math.Sqrt(7))
}

//与java不同 go的声明类型是放在变量后面
func add(x int, y int) int {
	return x + y;
}

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
func add1(x, y int) int {
	return x + y;
}

//go 函数可以返回任意数量的参数
func swap(x, y string) (string, string) {
	return y, x;
}

//Go 的返回值可被命名，它们会被当作定义在函数顶部的变量。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
