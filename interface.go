package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

// Rewriter writes to "destination" and counts bytes
type Rewriter struct {
	destination io.Writer
	count       int
}

func (r *Rewriter) Write(bytes []byte) (int, error) {
	r.count += len(bytes)
	return r.destination.Write(bytes)
}
func main() {
	r := Rewriter{destination: os.Stdout}
	fmt.Fprintf(&r, "test1\n")
	fmt.Fprintf(&r, "test2\n")
	fmt.Printf("Wrote %v bytes\n", r.count)
	math.Log2(6)
}
