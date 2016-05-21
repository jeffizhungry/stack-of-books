package facade

import "fmt"

type circle struct{}

func (c circle) draw() {
	fmt.Println("Drawing Circle")
}

type square struct{}

func (s square) draw() {
	fmt.Println("Drawing Square")
}

func DrawCircle() {
	c := circle{}
	c.draw()
}

func DrawSquare() {
	s := square{}
	s.draw()
}
