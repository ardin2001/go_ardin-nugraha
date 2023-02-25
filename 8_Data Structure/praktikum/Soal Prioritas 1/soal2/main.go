package main

import (
	"fmt"
)

func Mapping(arr []string) map[string]int {
	var container map[string]int
	container = map[string]int{}
	for _, data := range arr {
		stack := 0
		for _, data2 := range arr {
			if data == data2 {
				stack += 1
			}
		}
		container[data] = stack
	}
	return container
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
