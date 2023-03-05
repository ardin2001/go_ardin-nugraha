package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func Prints(print []pair) {
	for _, data := range print {
		fmt.Print(data.name, "->", data.count, " ")
	}
	fmt.Println()
}
func BubbleShort(arr []pair) []pair {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].count > arr[j+1].count {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	Prints(arr)
	return arr
}

func MostAppearItem(items []string) []pair {
	tempMap := make(map[string]int)
	for _, item := range items {
		tempMap[item] += 1
	}
	container := make([]pair, len(tempMap))
	i := 0
	for key, value := range tempMap {
		container[i] = pair{key, value}
		i++
	}
	return BubbleShort(container)
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tennis"}))
}
