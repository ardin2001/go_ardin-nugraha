package main

import (
	"fmt"
	"math"
)

func ArrayMulti(arr [][]int) []int {
	diagonal_kiri := 0
	diagonal_kanan := 0
	for i, data1 := range arr {
		for j := range arr {
			if i == j {
				diagonal_kiri += data1[i]
			}
			if i+j == len(arr)-1 {
				diagonal_kanan += arr[i][j]
			}
		}
	}
	return []int{diagonal_kiri, diagonal_kanan}
}
func main() {
	diagonal := ArrayMulti([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}})
	fmt.Println("Diagonal Kiri :", diagonal[0])
	fmt.Println("Diagonal Kanan :", diagonal[1])
	result := math.Abs(float64(diagonal[0] - diagonal[1]))
	fmt.Println("Perbedaan mutlak adalah", diagonal[0], "-", diagonal[1], "=", result)

}
