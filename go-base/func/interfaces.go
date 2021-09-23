package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type myf float64

type V struct {
	X, Y float64
}

func (f myf) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *V) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := myf(-math.Sqrt2)
	v := V{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())
}
