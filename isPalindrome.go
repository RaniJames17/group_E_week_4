package main

import "strings"

func isPalindrome(s string) bool {
	// Convert the string to lowercase and also remove spaces remove spaces for comparison
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	// Get the length of the string
	length := len(s)

	// Compare characters from the beginning and end
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}

	return true
}
