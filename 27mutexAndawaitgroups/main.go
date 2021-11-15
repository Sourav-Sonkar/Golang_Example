package main

import (
	"fmt"
	"sync"
)

//go run ---race .
//R - goRoutines
func main() {
	fmt.Println("Race conditions")

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("One R")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Two R")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Three R")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Reading scores by locking read value in  R")
		mut.RLock()
		fmt.Println("Score:", score)
		mut.RUnlock()
		wg.Done()

	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
