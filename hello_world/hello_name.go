package main

import "fmt"

func main() {
	var name string
	fmt.Println("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}
