// Factor pattern allows users to reate polymorphic types given a paramater
// such as a string describing the type.

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
	fmt.Println("Draw Circle")
}

type Square struct{}

func (s Square) draw() {
	fmt.Println("Draw Square")
}

/**********
 * Factory
 **********/

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return Circle{}
	case "SQUARE":
		return Square{}
	default:
		return nil
	}
}

/*******
 * Main
 *******/

func main() {
	shape1 := ShapeFactory("CIRCLE")
	shape1.draw()

	shape2 := ShapeFactory("SQUARE")
	shape2.draw()
}
