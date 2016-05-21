// Decorator pattern adds functionality to an existing class by wrapping the
// original class.
//
// The example below wraps the Circle class with the RedShape class, which
// both satisfies the Shape inteface.

package main

import "fmt"

/*************
 * Interface
 ************/

type Shape interface {
	draw()
}

/**********
 * Object
 *********/

type Circle struct{}

func (c Circle) draw() {
	fmt.Println("Drawing Circle\n")
}

/************
 * Decorator
 ************/

type RedShape struct {
	Shape
}

func NewRedShape(s Shape) RedShape {
	return RedShape{
		Shape: s,
	}
}

func (r RedShape) draw() {
	r.addBorder()
	r.Shape.draw()
}

func (r RedShape) addBorder() {
	fmt.Println("Adding Red Border")
}

/*******
 * Main
 *******/

func main() {
	var circleShape Shape
	circleShape = Circle{}
	circleShape.draw()

	var redCircleShape Shape
	redCircleShape = NewRedShape(circleShape)
	redCircleShape.draw()
}
