package main

import (
	"fmt"
	"time"
)

func Kelipatan(x int) {
	fmt.Println("Ready Go")
	for i := 1; true; i++ {

		fmt.Println(i * x)
		time.Sleep(3 * time.Second)
	}
}
func main() {
	go Kelipatan(2)
	time.Sleep(15 * time.Second)
	// gunakan select agar perulangan dilakukan terus menerus
	//select {}
}
