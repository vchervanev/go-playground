package main

import (
	"flag"
	"fmt"
)

var sep = flag.String("s", "default", "separator")

type point struct {
	X, Y int
}

var p point

func main() {
	p = point{X: 10, Y: 20}
	fmt.Println(*sep)
	fmt.Println(p)
}
