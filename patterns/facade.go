// Facade pattern hides complexities of the system by providing an interface
// for the client to access the system.
//
// Note: This is pretty intuitive, why does this need to be a pattern?

package main

import "./facade"

func main() {
	facade.DrawCircle()
	facade.DrawSquare()
}
