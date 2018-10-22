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

// Mover interface
type Mover interface {
	Move(dx, dy int)
}

// Player is a player
type Player struct {
	Point
	Name string
}

// NewPlayer is a function
func NewPlayer(name string, x, y int) *Player {
	return &Player{
		Point: Point{x, y},
		Name:  name,
	}
}

func moveAll(items []Mover, dx, dy int) {
	for _, m := range items {
		m.Move(dx, dy)
	}
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

	alice := NewPlayer("Alice", 10, 20)
	fmt.Printf("%+v\n", alice)

	ms := []Mover{&p1, &p2, alice}
	fmt.Println("ms before:", p1, p2, *alice)
	moveAll(ms, 20, 30)
	fmt.Println("ms after:", p1, p2, *alice)
}
