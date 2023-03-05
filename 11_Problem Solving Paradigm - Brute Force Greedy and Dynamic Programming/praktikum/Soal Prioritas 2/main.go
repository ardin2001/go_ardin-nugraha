package main

import "fmt"

func Minimum(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Frog(jump []int) int {
	dp := make([]int, len(jump))
	dp[0] = 0

	for i := 1; i < len(jump); i++ {
		dp[i] = Minimum(dp[i-1]+Absolute(jump[i]-jump[i-1]), dp[Maximum(0, i-2)]+Absolute(jump[i]-jump[Maximum(0, i-2)]))
	}

	return dp[len(jump)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
