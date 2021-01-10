package main

import (
	"fmt"
)

func typeOfInterface(s fmt.Stringer) {

	switch v := s.(type) {
	case Person:
		fmt.Printf("%v is a %T", v.String(), v)
	case Student:
		fmt.Printf("Student details: %v", v.String())
	default:
		fmt.Printf("Details: %v %T", v.String(), v)
	}
	fmt.Println()
}
