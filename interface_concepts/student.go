package main

import (
	"fmt"
)

type Student struct {
	name   string
	rollno int
}

/* Mutators */
func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetRollNo(rollno int) {
	s.rollno = rollno
}

/* Accessors*/

func (s *Student) Name() string {
	return s.name
}

func (s *Student) RollNo() int {
	return s.rollno
}

// implementing Stringers interface in fmt package
//A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
func (s Student) String() string {
	return fmt.Sprintf("%v (%v roll no)", s.Name(), s.RollNo())
}
