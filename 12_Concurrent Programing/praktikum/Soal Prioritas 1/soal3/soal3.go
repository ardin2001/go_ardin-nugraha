package main

import (
	"fmt"
	"time"
)

func main() {
	chn := make(chan int, 5)
	go func(c chan int) {
		// kapasitas 5 pada buffer channel tidak akan habis, karena buffer channel ketika dibuat selalu langsung dikeluarkan melalui perulangan
		for i := 1; true; i++ {
			c <- i * 3
			fmt.Println(<-c)
			time.Sleep(2 * time.Second)
		}
	}(chn)
	//time.Sleep(11 * time.Second)
	select {}
}
