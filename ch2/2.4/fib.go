package main

import "fmt"

func main() {
	// Tuple assignment in action!
	x, y := 1, 2

	fmt.Println("Fibanacci Sequence:")
	for i := 0; i < 10; i++ {
		fmt.Println(x)
		// THis is the magic line.
		// x gets the old valueof y.
		// y gets the old value of x + y.
		// It happens at the Exact same time!
		x, y = y, x+y
	}
}
