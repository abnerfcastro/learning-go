package main

import "fmt"

// Sample input
// 6
// 1 2 3 4 10 11

// Sample output
// 31

func sumArray(arr []int) int {
	var sum int
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

func main() {
	var length int
	fmt.Scan(&length)

	var array = make([]int, length)

	// read inputs
	for i := 0; i < length; i++ {
		fmt.Scan(&array[i])
	}

	// print sum
	fmt.Println(sumArray(array))
}
