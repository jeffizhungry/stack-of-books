// Filter pattern or Criteria pattern enables filtering objects using different
// criteria and chaining them in a decoupled manner.

package main

import "fmt"

/************************
 * Object to be filtered
 ************************/

type Person struct {
	Name          string
	Gender        string
	MartialStatus string
}

func NewPerson(name string, gender string, martialStatus string) Person {
	return Person{
		Name:          name,
		Gender:        gender,
		MartialStatus: martialStatus,
	}
}

/*******************
 * Filter interface
 *******************/

type Criteria interface {
	meetCriteria(pp []Person) []Person
}

/**********************
 * Filters / Criterias
 **********************/

type Male struct{}

func (c Male) meetCriteria(pp []Person) []Person {
	filt := make([]Person, 0, len(pp))
	for _, p := range pp {
		if p.Gender == "MALE" {
			filt = append(filt, p)
		}
	}
	return filt
}

type Female struct{}

func (f Female) meetCriteria(pp []Person) []Person {
	filt := make([]Person, 0, len(pp))
	for _, p := range pp {
		if p.Gender == "FEMALE" {
			filt = append(filt, p)
		}
	}
	return filt
}

type Single struct{}

func (s Single) meetCriteria(pp []Person) []Person {
	filt := make([]Person, 0, len(pp))
	for _, p := range pp {
		if p.MartialStatus == "SINGLE" {
			filt = append(filt, p)
		}
	}
	return filt
}

/**********************************
 * Filters / Criterias Combinators
 **********************************/
// This kinda uses a decorator pattern as well

type Or struct {
	C1 Criteria
	C2 Criteria
}

func (o Or) meetCriteria(pp []Person) []Person {
	first := o.C1.meetCriteria(pp)
	second := o.C2.meetCriteria(pp)

	for _, sp := range second {
		var found bool
		for _, fp := range first {
			if fp.Name == sp.Name {
				found = true
			}
		}
		if !found {
			first = append(first, sp)
		}
	}
	return first
}

type And struct {
	C1 Criteria
	C2 Criteria
}

func (a And) meetCriteria(pp []Person) []Person {
	return a.C1.meetCriteria(a.C2.meetCriteria(pp))
}

/*******
 * Main
 *******/

func PrintPersons(pp []Person) {
	for _, p := range pp {
		fmt.Printf("%+v\n", p)
	}
}

func main() {
	ppl := []Person{
		NewPerson("Robert", "MALE", "SINGLE"),
		NewPerson("John", "MALE", "MARRIED"),
		NewPerson("Laura", "FEMALE", "SINGLE"),
		NewPerson("Diana", "FEMALE", "MARRIED"),
		NewPerson("Mike", "MALE", "SINGLE"),
		NewPerson("Bobby", "MALE", "SINGLE"),
	}

	var male Criteria = Male{}
	var female Criteria = Female{}
	var singleMale Criteria = And{Male{}, Single{}}
	var singleOrFemale Criteria = Or{Female{}, Single{}}

	fmt.Println("Males:")
	PrintPersons(male.meetCriteria(ppl))
	fmt.Println("\nFemales:")
	PrintPersons(female.meetCriteria(ppl))
	fmt.Println("\nSingle Males:")
	PrintPersons(singleMale.meetCriteria(ppl))
	fmt.Println("\nSingle Or Females:")
	PrintPersons(singleOrFemale.meetCriteria(ppl))
}
