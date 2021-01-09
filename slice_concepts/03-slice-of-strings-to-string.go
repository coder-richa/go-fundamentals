package main

import (
	"fmt"
	"strings"
)

func joinElementsOfSlice() {
	// Declare and initialise a slice of string
	fruits := []string{"Apple", "Mango", "Orange", "Peach", "Papaya"}
	fmt.Println("Slice:", fruits)
	// Joining slice elements with ',' separator
	output := strings.Join(fruits, ", ")
	fmt.Println(output)
	/*
		Output:
		Slice: [Apple Mango Orange Peach Papaya]
		Apple, Mango, Orange, Peach, Papaya
	*/
}
