package main

import (
	"fmt"
)

func CekData(arr []string) map[string]int {
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
	fmt.Println(CekData([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(CekData([]string{"ardin", "nugraha", "ardin"}))
	fmt.Println(CekData([]string{}))
}
