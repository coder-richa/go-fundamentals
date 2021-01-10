package main

import (
	"fmt"
)

func main() {
	p := &Person{}
	p.SetName("John")
	p.SetAge(18)
	fmt.Println("Person details:", p.String())
	/*
		Output:
		Person details: John (18 years)
	*/

	anonymousStruct()
	anonymousStructSlice()
}
