package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Group E's Week 4 Project!")

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

	// calling factorial for 12 ,

	string_num := strconv.Itoa(factorial(12))
	fmt.Println("The factorial of 12 is " + string_num)

	// checking if first string entered is  palindrome
	result := isPalindrome(input)
	fmt.Printf("The given string is a palindrome is a %t statement\n", result)
}

// Function to check if a string contains a substring
func containsString(str string) bool {
	var substr string

	fmt.Print("Enter a substring:")
	fmt.Scanf("%s", &substr)
	return strings.Contains(str, substr)
}
