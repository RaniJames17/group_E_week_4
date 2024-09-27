package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to check if a string contains a substring
func containsString(str string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a substring: ")
	substr, _ := reader.ReadString('\n')
	substr = strings.TrimSpace(substr)
	return strings.Contains(str, substr)
}
