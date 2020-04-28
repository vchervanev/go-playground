package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// for i := 0; i < len(os.Args); i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	fmt.Println(os.Args[i])
	// }

	// for i := 5; i < 500; {
	// 	i += 150
	// 	fmt.Println(i)
	// }

	// for a, b := range os.Args {
	// 	fmt.Println(a, b)
	// }
	fmt.Println(os.Args)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := strings.ReplaceAll(input.Text(), "a", "\033[1;33ma\033[0m")
		fmt.Println(t)
	}

	index := make(map[string]int)
	index["a"] = 100
	index["b"] = 200

	for key, value := range index {
		fmt.Print(key)
		fmt.Print("=>")
		fmt.Println(value)
	}
}
