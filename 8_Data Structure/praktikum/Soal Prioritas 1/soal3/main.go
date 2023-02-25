package main

import (
	"fmt"
)

func MunculSekali(arr string) []string {
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
	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
}
