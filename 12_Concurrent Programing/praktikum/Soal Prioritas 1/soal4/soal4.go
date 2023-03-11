package main

import (
	"fmt"
	"sync"
)

type SafeNumber struct {
	i int
	m sync.Mutex
}

func (safeNumber *SafeNumber) setNumber(number int) {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()
	safeNumber.i = number
}
func (safeNumber *SafeNumber) GetNumber() int {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()
	result := 1
	for i := 2; i <= safeNumber.i; i++ {
		result *= i
	}
	return result
}
func faktorial(number int) int {
	var wg sync.WaitGroup
	wg.Add(1)
	var i = SafeNumber{}
	go func() {
		i.setNumber(number)
		wg.Done()
	}()
	wg.Wait()
	return i.GetNumber()
}
func main() {
	fmt.Println(faktorial(4))
}
