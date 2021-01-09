package main

import (
	"fmt"
)

func removeElementFromSlice() {
	// Declare and initialise a slice of string
	fruits := []string{"Apple", "Mango", "Orange", "Peach", "Papaya"}
	fmt.Println("Before removing elements:", fruits)
	// removing elements to the slice
	fruits = append(fruits[:1], fruits[3:]...)
	fmt.Println("After removing elements:", fruits)
	/*
		Output:
		Before removing elements: [Apple Mango Orange Peach Papaya]
		After removing elements: [Apple Peach Papaya]
	*/
}
