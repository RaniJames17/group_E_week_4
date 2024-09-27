package main

import (
	"fmt"
	"math"
)

func calculateSquareRoot() {
	var a float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	if a >= 0 {
		result := math.Sqrt(a)
		fmt.Printf("The square root of %.2f is %.2f\n", a, result)
	} else {
		fmt.Println("Please enter a non-negative number.")
	}
}
