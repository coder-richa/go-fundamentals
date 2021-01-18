//https://www.devdungeon.com/content/working-files-go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getFileInfo(path string) {
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func writeDataToFile(path string, byteSlice []byte) (err error, success bool, bytesWritten int) {
	// Open a new file for writing only
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err, success, bytesWritten
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err = file.Write(byteSlice)
	return err, err == nil, bytesWritten
}

func readDataFromFile(path string) (err error, success bool, data []byte) {
	// Open file for reading
	file, err := os.Open(path)
	if err != nil {
		return err, success, data
	}

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// ioutil.ReadAll() will read every byte
	// from the reader (in this case a file),
	// and return a slice of unknown slice
	data, err = ioutil.ReadAll(file)
	return err, err == nil, data
}
