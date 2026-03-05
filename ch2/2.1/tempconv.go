package main

import "fmt"

// Package -level constant: visible to the whole file
const boilingF = 212.0

func main() {
	// Local variables: only visible inside main()
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
