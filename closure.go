package main

import "fmt"

func multiplyBy(x int) func(int) int {
	return func(v int) int {
		return x * v
	}
}

func main() {
	f1 := multiplyBy(5)
	f2 := multiplyBy(7)
	fmt.Println(f1(5))
	fmt.Println(f2(3))
}
