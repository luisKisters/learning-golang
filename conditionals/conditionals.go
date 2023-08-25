package main

import "fmt"

func main() {
	fmt.Println("Do you want to hear a joke? (y/n)")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" {
		fmt.Println("What do you call a fat psychic? A four-chin teller.") // joke from https://icanhazdadjoke.com/
	}
	if answer == "n" {
		fmt.Println("Okay:(")
	}
}
