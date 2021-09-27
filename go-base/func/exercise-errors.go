package main

import "fmt"

type ErrNegativeSqrt float64

func (t ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("canot Sqrt negative number:%v", float64(t))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
