package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10000 {
		fmt.Print("\r", i)
		time.Sleep(time.Millisecond)
		i++
	}
}
