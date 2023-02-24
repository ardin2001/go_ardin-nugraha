package main

import (
	"fmt"
)

func kuadrat(x, y int) int {
	value := 1
	i := 1
	for y >= i {
		if y%2 == 0 {
			value = value * x * x
			y -= 2
		} else if y%2 != 0 {
			value = value * x
			y--
		}
	}
	return value
}
func main() {
	fmt.Println(kuadrat(2, 16))

}
