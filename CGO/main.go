package main

import "C"
import "unsafe"

func main()  {
	v := 42
	C.printint(C.int(v))
}