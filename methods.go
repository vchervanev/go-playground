package main

import "fmt"

type point struct {
	X int
	Y int
}

type rect struct {
	P1 point
	P2 point
}

func (p1 point) rectangle(p2 point) rect {
	return rect{P1: p1, P2: p2}
}

func (r rect) dimensions() (dx, dy int) {
	return r.dx(), r.dy()
}

func (r rect) dx() int {
	return r.P2.X - r.P1.X
}

func (r rect) dy() int {
	return r.P2.Y - r.P1.Y
}

func (r rect) area() int {
	dx, dy := r.dimensions()
	return dx * dy
}

func main() {
	p1 := point{10, 20}
	p2 := point{15, 30}
	fmt.Printf("[%#v-%#v] area = %d\n", p1, p2, p1.rectangle(p2).area())
}
