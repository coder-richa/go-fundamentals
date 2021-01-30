/**************************************************************************

Type assertions

A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)

If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

*/
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
	// Type Assertion
	y,ok:=x.(*Teacher)
	if !ok{
		fmt.Println("Error occured while type assertion")
	}	
	fmt.Printf("%T Details  %v",y,y)
	/************
	Output
	*main.Teacher Details  &{Richa 18}
	*/
	
}
