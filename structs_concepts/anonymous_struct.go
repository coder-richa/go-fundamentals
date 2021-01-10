package main

import (
	"fmt"
)

func anonymousStruct() {
	p := struct {
		name string
		age  int
	}{
		name: "John",
		age:  18,
	}
	fmt.Println("Anonymous struct details:", p)
	/*
		Output:
		Anonymous struct details: {John 18}
	*/
}

func anonymousStructSlice() {
	p := []struct {
		name string
		age  int
	}{
		{
			name: "John",
			age:  18,
		}, {
			name: "Jane",
			age:  19,
		},
	}
	fmt.Println("Anonymous slice of struct details:", p)
	/*
		Output:
		Anonymous slice of struct details: [{John 18} {Jane 19}]
	*/
}
