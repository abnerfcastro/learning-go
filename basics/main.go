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

	// for loops
	for i := 0; i < 10; i++ {
		fmt.Printf("Looping in binary: %d - %b\n", i, i)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("Using variables from other packages")
	fmt.Println(utils.MyName)
	fmt.Println(utils.WhoAmI())
	fmt.Printf("Using function Sum: %d\n", utils.Sum(1, 2))

	// variables
	var a = "initial" // type will be infered
	var b, c int = 1, 2
	fmt.Println(a, b, c)

	// := syntax is shorthand for declaring and initializing a variable
	f := "short way"
	fmt.Println(f)

	// constants
	const n = 50

	// There is no ternary if in Go
	// so you’ll need to use a full if statement even for basic conditions

	// approach to ternary
	ternary := map[bool]int{true: 1, false: 0}[5 > 4]
	fmt.Println("ternary:", ternary)

	// if conditions
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// arrays
	var arr [5]int
	fmt.Println("emp:", arr)

	fmt.Println("len:", len(arr))

	integers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", integers)

}
