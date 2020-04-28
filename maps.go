package main

import "fmt"

func main() {
	m := map[string]int{"a": 42}

	fmt.Println("a", m["a"])
	fmt.Println("b", m["b"])
}
