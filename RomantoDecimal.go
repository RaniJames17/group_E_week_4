package main

import (
	"fmt"
	"strings"
)

var romanToDecimal = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(roman string) {
	roman = strings.ToUpper(roman)
	length := len(roman)
	result := 0

	for i := 0; i < length; i++ {
		if i+1 < length && romanToDecimal[roman[i]] < romanToDecimal[roman[i+1]] {
			result -= romanToDecimal[roman[i]]
		} else {
			result += romanToDecimal[roman[i]]
		}
	}

	fmt.Printf("The decimal value of %s is %d\n", roman, result)

}
