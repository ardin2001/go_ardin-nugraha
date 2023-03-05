package main

import (
	"fmt"
)

func PlayingDomino(card [][]int, deck []int) interface{} {
	for _, data := range card {
		for _, data2 := range deck {
			if data[0] == data2 || data[1] == data2 {
				return data
			}
		}
	}
	return "Tutup Kartu"
}

func main() {
	fmt.Println(PlayingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(PlayingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(PlayingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
