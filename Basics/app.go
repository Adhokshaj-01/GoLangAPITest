package main

import "fmt"

func main() {
	// sum()
	// array()
	arr := []int{10, 20, 30, 40, 50}

	var target int

	// Input the target number to search for
	fmt.Print("Enter the number to search: ")
	fmt.Scanln(&target)

	// Call linearSearch function
	result := linearSearch(arr, target)

	if result != -11 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found.")
	}
}

// add two numbers
//define parameters and return type
func sum() {
	var var1, var2 int
	fmt.Println("Hellow From GO lang")

	fmt.Println("Enter First Number :")
	fmt.Scanln(&var1)

	fmt.Println("Enter Second Number :")
	fmt.Scanln(&var2)

	sum := var1 + var2

	fmt.Print("The Sum is ", sum)
}

// initilizing array
func array() {
	// Array of 5 integers
	var arr = [5]int{1, 2, 3, 4, 5}
	// The size is inferred by the number of elements
	arr1 := [...]int{1, 2, 3, 54, 6, 7, 7, 8, 8, 8}
	// Partial initialization, rest will be default values (zero for int)
	arr2 := [5]int{0: 1, 3: 50}
	// Print the array
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

}

// linear search in go
func linearSearch(arr []int, target int) int {
	for index, e := range arr {
		if e == target {
			return index
		}
	}
	return -11
}
