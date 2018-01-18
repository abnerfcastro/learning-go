package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program name: %s\n", os.Args[0])
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
