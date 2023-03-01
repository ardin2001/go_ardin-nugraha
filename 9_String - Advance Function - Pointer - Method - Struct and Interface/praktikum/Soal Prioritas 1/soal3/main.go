package main

import (
	"fmt"
)

func Compare(a, b string) string {
	result := ""
	for i, data1 := range a {
		for j, data2 := range b {
			if string(data1) == string(data2) && i == j {
				result += string(data2)
			}

		}
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
