package main

import (
	"fmt"
)

func PairSum(arr []int, target int) []int {
	result := make([]int, 0)
outerLoop:
	for i, data1 := range arr {
		for j, data2 := range arr {
			if data1+data2 == target {
				result = append(result, i)
				result = append(result, j)
				break outerLoop
			}
		}
	}
	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
