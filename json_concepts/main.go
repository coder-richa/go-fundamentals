package main

import (
	"fmt"
)

func main() {
	p := &Person{}
	p.SetName("Raman")
	p.SetAge(18)
	b, err := jsonEncode(p)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Encoded string", string(b))
		// Initialise variable to hold json data
		x := &Person{}
		err = jsonDecode(b, x)
		if err != nil {
			fmt.Println("Error", err.Error())
		} else {
			fmt.Println("Decoded Person", x)
		}
	}
}
