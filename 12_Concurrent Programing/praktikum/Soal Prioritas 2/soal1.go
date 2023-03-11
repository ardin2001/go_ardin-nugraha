package main

import (
	"fmt"
	"sync"
)

func main() {
	teks := "helloworld"
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(map[string]int)
	go func() {
		left := teks[:len(teks)/2]
		for _, data := range left {
			result[string(data)] += 1
		}
		wg.Done()
	}()
	go func() {
		right := teks[len(teks)/2:]
		for _, data := range right {
			result[string(data)] += 1
		}
		wg.Done()
	}()
	wg.Wait()
	for key, val := range result {
		fmt.Println(key, ":", val)
	}
}
