package main

// factorial fucntion using recursion
func factorial(a int) int {
	if a == 0 || a == 1 {
		// as 1 and zero's factorial is 1 so we just return 1
		return a
	}

	// in other cases we keep calling factorial until the argument is reduced to 1
	return a * factorial(a-1)
}
