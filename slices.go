package main

import "fmt"

func swap(s []int, a int, b int) {
	s[a], s[b] = s[b], s[a]
}

func main() {
	input := [...]int{10, 20, 30, 40, 50}
	fmt.Println("Input", input)
	s := input[1:4]
	fmt.Println("Slice", s)
	swap(s, 0, 2)
	fmt.Println("Transformed Slice", s)
	fmt.Println("Transformed Input", input)
}
