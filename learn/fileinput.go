package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// inputFile, err := os.Open("test.go")
	// if err != nil {
	// 	fmt.Println("12 err:", err)
	// 	return
	// }
	// defer inputFile.Close()

	// inputReader := bufio.NewReader(inputFile)
	// for {
	// 	inputString, readErr := inputReader.ReadString('\n')
	// 	fmt.Printf("The input was: %s", inputString)
	// 	if readErr == io.EOF {
	// 		return
	// 	}
	// }
	together()
}

func together() {
	inputFile := "test.go"
	outputFile := "test_copy.go"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Println("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
