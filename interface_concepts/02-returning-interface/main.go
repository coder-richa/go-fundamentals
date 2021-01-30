package main

import (
	"fmt"
)

type Teacher struct{
	name string
	age int
}

func (t Teacher) details() (interface{}) {
	// Returns Pointer to Teacher type
	return &t
}

type Print interface{
	details() interface{}	
}

func main() {
	t:=Teacher{"Richa",18}
	x:=t.details()	
	switch v := x.(type){
	 case *Teacher:
		// Executes when it receives Pointer to Teacher
		fmt.Printf("%T Details  %v",v,x)
	 default:
		fmt.Printf("%T",v)
	}
	/************
	Output
	*main.Teacher Details  &{Richa 18}
	*/	
}
