package main

import (
	"fmt"
)

func CekData(arr string) []string {
	container := make([]string, 0)
	for _, data := range arr {
		stack := 0
		for _, data2 := range arr {
			if data == data2 {
				stack += 1
			}
		}
		if stack == 1 {
			container = append(container, string(data))
		}
	}
	return container
}

func main() {
	fmt.Println(CekData("1234123"))
	fmt.Println(CekData("76523752"))
	fmt.Println(CekData("12345"))
	fmt.Println(CekData("1122334455"))
	fmt.Println(CekData("0872504"))
}
