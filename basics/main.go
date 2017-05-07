package main

import (
	"fmt"

	"github.com/abnerfcastro/learning-go/basics/utils"
)

func main() {
	// Hello World must :)
	fmt.Println("Hello World!")

	// Printing formats
	fmt.Printf("Printing decimal: %d\n", 23)
	fmt.Printf("Printing binary: %d - %b\n", 23, 23)
	fmt.Printf("Printing hexadecimal: %d \t %b \t %#X \n", 23, 23, 23)

	// For loops
	for i := 0; i < 10; i++ {
		fmt.Printf("Looping in binary: %d - %b\n", i, i)
	}

	fmt.Println("Using variables from other packages")
	fmt.Println(utils.MyName)
	fmt.Println(utils.WhoAmI())
	fmt.Printf("Using function Sum: %d\n", utils.Sum(1, 2))
}
