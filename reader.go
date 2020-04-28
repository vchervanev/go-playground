package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		str, err := in.ReadString('\n')

		if err == io.EOF {
			return
		}
		fmt.Println(">>", str)
	}
}
