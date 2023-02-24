package main

import "fmt"

func Result(merge []string) []string {
	concats := make([]string, 0)
	for _, data := range merge {
		stack := 0
		for _, concat := range concats {
			if data != concat {
				stack += 1
			}
		}
		if stack == len(concats) {
			concats = append(concats, data)
		}
	}
	return concats

}
func ArrayMerge(arrayA, arrayB []string) []string {
	gabungan := make([]string, 0, 5)
	gabungan = append(gabungan, arrayA...)
	gabungan = append(gabungan, arrayB...)
	return Result(gabungan)

}
func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
