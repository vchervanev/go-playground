package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mu sync.Mutex
var wg sync.WaitGroup

func print(id int, times int) {
	defer wg.Done()
	for ; times > 0; times-- {
		// mu.Lock()
		counter++
		// mu.Unlock()
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	wg.Add(50000)
	go spinner(100 * time.Millisecond)
	for i := 0; i < 50000; i++ {
		go print(i, 50000)
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
