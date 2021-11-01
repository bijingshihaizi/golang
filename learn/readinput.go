package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	// input                  = "56.12 / 5211 / go"
	format      = "%f / %d / %s"
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	// echo()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("the input was : %s\n", input)
	}
}

func echo() {
	fmt.Println("please enter ur full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
