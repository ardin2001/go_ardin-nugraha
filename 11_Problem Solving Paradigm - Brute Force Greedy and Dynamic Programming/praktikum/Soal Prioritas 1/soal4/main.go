package main

import (
	"fmt"
	"math"
)

func solveEquations(a, b, c int) {
	x := int(math.Cbrt(float64(b)))
	if a-x < 0 {
		fmt.Println("No Solution")
		return
	} else if a-x == 0 {
		fmt.Println("No Solution")
		return
	}

	z := int(math.Sqrt(float64(c - x*x)))
	if a-x-z < 0 {
		fmt.Println("No Solution")
		return
	} else if a-x-z == 0 {
		fmt.Println("No Solution")
		return
	}

	y := a - x - z
	if x*x+y*y+z*z != c || y*z*x != b {
		fmt.Println("No Solution")
	}

	fmt.Println(x, y, z)
}

func main() {
	solveEquations(1, 2, 3)
	solveEquations(6, 6, 14)
}
