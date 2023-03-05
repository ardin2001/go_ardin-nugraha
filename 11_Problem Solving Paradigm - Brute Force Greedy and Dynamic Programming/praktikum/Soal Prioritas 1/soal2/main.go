package main

import "fmt"

func SegitigaPascal(rows int, base int) [][]int {
	containers := make([][]int, 0)
	for i := 0; i < rows; i++ {
		container := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || i == 0 {
				base = 1
				container = append(container, base)
			} else {
				base = base * (i - j + 1) / j
				container = append(container, base)
			}
		}
		containers = append(containers, container)
	}
	return containers
}
func main() {
	var rows int
	fmt.Print("Input : ")
	fmt.Scan(&rows)
	fmt.Println("Output :", SegitigaPascal(rows, 1))

}
