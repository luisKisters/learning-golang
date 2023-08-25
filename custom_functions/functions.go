package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var num1 int
	var num2 int
	fmt.Println("Welcome to my calculator!")
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Println("The sum of ", num1, " and ", num2, " is ", add(num1, num2))
}
