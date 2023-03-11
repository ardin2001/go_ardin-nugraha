package main

import (
	"fmt"
	"time"
)

func Kelipatan(x chan int) {
	value := <-x
	fmt.Println("Ready Go")
	for i := 1; true; i++ {

		fmt.Println(i * value)
		time.Sleep(3 * time.Second)
	}
}
func main() {
	chn := make(chan int)
	go Kelipatan(chn)
	chn <- 3
	time.Sleep(15 * time.Second)
	// gunakan select agar perulangan dilakukan terus menerus
	//select {}
}
