package main

import "fmt"

func main() {
	var stringVar string
	var intVar int
	var boolVar bool

	fmt.Println("Input a string")
	fmt.Scan(&stringVar)
	fmt.Println("Input a integer")
	fmt.Scan(&intVar)
	fmt.Println("Input a boolean ('true' or 'false')")
	fmt.Scan(&boolVar)

	fmt.Println("the stringVar you set is: ", stringVar)
	fmt.Println("the intVar you set is: ", intVar)
	fmt.Println("the boolVar you set is: ", boolVar)
}
