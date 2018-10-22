package main

import (
	"fmt"
)

// Point is 2D point
type Point struct {
	X int
	Y int
}

// Move doc
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p1 := Point{1, 2}
	fmt.Println("p1", p1)
	fmt.Printf("x is %d\n", p1.X)

	p2 := Point{
		Y: 20,
		X: 10,
	}

	fmt.Println("p2", p2)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%+v\n", p2)
	fmt.Printf("%#v\n", p2)

	p1.Move(10, 20)
	fmt.Printf("p1 after move: %+v\n", p1)
}
