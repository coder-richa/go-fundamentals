package main

import (
	"fmt"
)

func addElementsToSlice() {
	// Declare and initialise a slice of string
	fruits := []string{"Apple", "Mango"}
	fmt.Println("Before adding elements:", fruits)
	// adding elements to the slice
	fruits = append(fruits, "Orange", "Peach", "Papaya")
	fmt.Println("After adding elements:", fruits)
	/*
		Output:
		Before adding elements: [Apple Mango]
		After adding elements: [Apple Mango Orange Peach Papaya]
	*/
}
