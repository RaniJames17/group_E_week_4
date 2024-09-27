package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: \n")

	// Read user input
	input, _ := reader.ReadString('\n')

	// Trim the newline Character from input
	input = strings.TrimSpace(input)

	//the result is dispalyed
	contains := containsString(input)
	if !contains {
		fmt.Println("Substring does not exist")
	} else {
		fmt.Printf("Substring exists")
	}
}

// Function to check if a string contains a substring
func containsString(str string) bool {
	var substr string

	fmt.Print("Enter a substring:")
	fmt.Scanf("%s", &substr)
	return strings.Contains(str, substr)
}
