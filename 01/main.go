package main

import (
	"fmt"
	"golang-learning/variables"
)

func main() {

	// short declaration (most common inside functions)
	// greeting := "Hello, World!"
	// age := 20

	// Function Declaration
	plus := func(a, b int) int {
		return a + b
	}
	pi := 3.14

	// Function Call
	result := plus(2, 5)

	// Print Statement
	fmt.Println(variables.Message, variables.Age, result)

	// float variable declaration
	fmt.Println(pi)
	// float32 vs float64
	// float32 is more precise than float64
	// float32 is faster than float64
	// float32 is more memory efficient than float64
	// float32 is more accurate than float64
	// float32 is more precise than float64
	// float32 is more accurate than float64
	pi2 := float64(pi)
	fmt.Println(pi2)

	// boolean variable declaration
	isStudent := false
	fmt.Println(isStudent)

	// rune and byte
	// rune is a single character

	var letter rune = '&'
	fmt.Println(letter, string(letter))
	var byte byte = 'A'
	fmt.Println(byte, string(byte))

	fullName := "Mohammad Garmabi"

	fmt.Printf("Full Name: %v , Age: %v - typed: %T", fullName, variables.Age, variables.Age)
}
