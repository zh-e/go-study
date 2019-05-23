package main

import "fmt"

//变量命名相关

var c, py, java bool
var a, b = true, "no"


func main() {
	var i int
	fmt.Println(i, c, py, java)

	var j = 2
	fmt.Println(j, a, b)

	//:= 可在类型明确的地方代替 var 声明 但是在函数外不可使用
	k := 3
	fmt.Println(k)

	defaultValue()

	//常量的声明 常量不能用 := 赋值
	const name = "go"

}

func defaultValue() {
	var i int //0
	var f float64 //0
	var b bool //false
	var s string //""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
/* 基本类型有
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
// 表示一个 Unicode 码点

float32 float64

complex64 complex128

 */
