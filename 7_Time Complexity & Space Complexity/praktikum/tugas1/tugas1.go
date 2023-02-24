package main

import (
	"fmt"
)

func Prima(a int) string {
	if a == 2 || a == 3 || a == 5 || a == 7 {
		return "Bilangan Prima"
	}
	if a%2 == 0 || a%3 == 0 || a%5 == 0 || a%7 == 0 {
		return "Bukan bilangan prima"
	}
	return "Bilangan Prima"
}
func main() {
	fmt.Println(Prima(49))
	fmt.Println(Prima(1000000007))
	fmt.Println(Prima(1500450271))
	fmt.Println(Prima(1000000007))
	fmt.Println(Prima(13))
	fmt.Println(Prima(17))
	fmt.Println(Prima(20))
	fmt.Println(Prima(35))

}
