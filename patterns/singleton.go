// Singleton pattern makes sure that there is only one instance of an object
// that is ever created.
//
// Note: Particularly in Go, we have to create a package level singleton.

package main

import (
	"fmt"

	"./singleton"
)

func main() {
	obj1 := singleton.Instance()
	fmt.Printf("Singleton Object1 Addr = %p\n", obj1)

	obj2 := singleton.Instance()
	fmt.Printf("Singleton Object2 Addr = %p\n", obj2)
}
