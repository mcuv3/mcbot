package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	x := GMin(2, 3)
	fmt.Println(x)
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
