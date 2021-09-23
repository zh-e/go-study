package main

import "fmt"

type V2 struct {
	Lat, Long float64
}

var m map[string]V2

func main() {
	m = make(map[string]V2)
	m["Bell Labs"] = V2{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])

	var c = map[string]V2{
		"11": {
			34, 45,
		},
		"22": {
			20, 23,
		},
	}

	fmt.Println(c)

	v , ok := c["11"]
	fmt.Println(ok, v)

}
