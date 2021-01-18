package main

import (
	"fmt"
)

func main() {
	fileName := "person.txt"
	// getFileInfo("person.go")
	p := &Person{}
	p.SetName("Raman")
	p.SetAge(18)

	q := &Person{}
	q.SetName("Ram")
	q.SetAge(21)

	s := []*Person{p, q}

	b, err := jsonEncode(s)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		// write data to file
		err, success, bytesWritten := writeDataToFile(fileName, b)
		fmt.Println("error", err)
		fmt.Println("success", success)
		fmt.Println("bytesWritten", bytesWritten)

	}
	x := []*Person{}
	err, success, data := readDataFromFile(fileName)
	if success {
		err = jsonDecode(data, &x)
	}
	fmt.Println("error", err)
	fmt.Println("success", success)
	fmt.Println("bytes read", string(data))
	fmt.Println("Person from File", x)
	/*
		Output:
		error <nil>
		success true
		bytesWritten 51
		error <nil>
		success true
		bytes read [{"name":"Raman","age":18},{"name":"Ram","age":21}]
		Person from File [Raman (18 years) Ram (21 years)]
	*/
}
