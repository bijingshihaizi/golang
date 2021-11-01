package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "this is an apple this"
	fmt.Println("t or f ", strings.HasPrefix(str, "this"))
	fmt.Println("t or f ", strings.Contains(str, "this"))
	fmt.Println("t or f ", strings.Index(str, "this"))
	fmt.Println("t or f ", strings.LastIndex(str, "this"))
	fmt.Println("t or f ", strings.Replace(str, "this", "that", 1))
	fmt.Println("t or f ", strings.Count(str, "this"))

}
