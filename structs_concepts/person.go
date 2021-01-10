package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

/* Mutators */
func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

/* Accessors*/

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

// implementing Stringers interface in fmt package
//A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name(), p.Age())
}
