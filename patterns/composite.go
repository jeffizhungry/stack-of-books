// Composite pattern treats a group of objects the same way as a singlular object
//
// Note: I don't know how useful this is... This is kinda stupid

package main

import "fmt"

type Employee struct {
	Name         string
	Subordinates []*Employee
}

func NewEmployee(name string) *Employee {
	return &Employee{
		Name:         name,
		Subordinates: make([]*Employee, 0, 0),
	}
}

func (e *Employee) Add(n *Employee) {
	e.Subordinates = append(e.Subordinates, n)
}

func (e *Employee) PrintTree(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("    ")
	}
	fmt.Println(e.Name)
	for _, s := range e.Subordinates {
		s.PrintTree(depth + 1)
	}
}

func main() {
	ceo := NewEmployee("John the CEO")
	leadSales := NewEmployee("Robert the Sales Lead")
	clerk1 := NewEmployee("Tammy the Sales Clerk")
	clerk2 := NewEmployee("Justing the Sales Clerk")
	leadMarket := NewEmployee("Jules the Marketing Lead")
	mark1 := NewEmployee("Tammy the Marketer")
	mark2 := NewEmployee("Justing the Marketer")

	ceo.Add(leadSales)
	ceo.Add(leadMarket)
	leadSales.Add(clerk1)
	leadSales.Add(clerk2)
	leadMarket.Add(mark1)
	leadMarket.Add(mark2)

	ceo.PrintTree(0)
}
