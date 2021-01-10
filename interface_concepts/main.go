package main

import (
	"fmt"
)

func main() {
	// Person Type
	p := Person{}
	p.SetName("John")
	p.SetAge(18)
	// Student Type
	s := Student{}
	s.SetName("Peter")
	s.SetRollNo(101)
	// *Person Type
	x := &Person{}
	x.SetName("John")
	x.SetAge(18)

	typeOfInterface(p)
	typeOfInterface(s)
	typeOfInterface(x)
	/*
		Output:
		John (18 years) is a main.Person
		Student details: Peter (101 roll no)
		Details: John (18 years) *main.Person
	*/

	c := &Circle{}
	c.SetRadius(10)
	fmt.Println("Area:", c.Area())
	fmt.Println("Circumference:", c.Perimeter())
	/*
		Output:
		Area: 314.1592653589793
		Circumference: 62.83185307179586
	*/
}
