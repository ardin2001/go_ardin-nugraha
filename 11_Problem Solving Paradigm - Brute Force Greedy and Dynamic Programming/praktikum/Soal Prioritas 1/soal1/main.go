package main

import (
	"fmt"
	"strconv"
)

func Biner(number int) []int {
	container := make([]int, 0)
	for i := 0; i <= number; i++ {
		if i == 0 {
			container = append(container, i)
			continue
		}
		binary := ""
		j := i
		for j > 0 {
			rem := j % 2
			binary = strconv.Itoa(rem) + binary
			j = j / 2
		}
		convertBinary, _ := strconv.Atoi(binary)
		container = append(container, convertBinary)
	}
	return container
}
func main() {
	fmt.Println(Biner(2))
	fmt.Println(Biner(3))
}
