package main

import "fmt"

// 1. Declare two distinct types that share the same underlying type (float64)
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	var c Celsius = FreezingC
	var f Fahrenheit = 32

	// fmt.Println(c == f)
	// ^ If you un-comment the line above, the code will CRASH.
	// Go will say: "mismatched types Celsius and Fahrenheit"

	// 2. To compare or combine them, you MUST explicitly convert the type
	convertedToF := Fahrenheit(c*9/5 + 32)

	fmt.Printf("%g°C is equal to %g°F\n", c, convertedToF)
}
