package main

import (
	"fmt"
)

type Person struct {
	PersonName string `json:"name"`
	PersonAge  int    `json:"age"`
}

/* Mutators */
func (p *Person) SetName(name string) {
	p.PersonName = name
}

func (p *Person) SetAge(age int) {
	p.PersonAge = age
}

/* Accessors*/

func (p *Person) Name() string {
	return p.PersonName
}

func (p *Person) Age() int {
	return p.PersonAge
}

// implementing Stringers interface in fmt package
//A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name(), p.Age())
}
