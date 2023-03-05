package main

import (
	"fmt"
	"math"
)

func Prima(number int) int {
	container := make([]int, 0)
	for i := 2; true; i++ {
		if len(container) == number {
			break
		}
		akarKuadrat := math.Sqrt(float64(i))
		if number < 1 {
			return -1
		} else if akarKuadrat >= 2 {
			for j := 2; j < i; j++ {
				stack := 0
				if i%j == 0 {
					break
				} else {
					stack += 1
				}
				if stack == i-j {
					container = append(container, i)
				}
			}
		} else {
			container = append(container, i)
		}
	}
	return container[number-1]
}
func main() {
	fmt.Println(Prima(1))
	fmt.Println(Prima(5))
	fmt.Println(Prima(8))
	fmt.Println(Prima(9))
	fmt.Println(Prima(10))

}
