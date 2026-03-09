package main

import "fmt"

func main() {
	// 1. Short Declaration (Concept 2.3.1)
	x := 1

	// 2. Pointers (Concept 2.3.2)
	p := &x         // p hold the address of x
	fmt.Println(*p) // Print1
	*p = 2          // Changes x to 2 through the pointer
	fmt.Println(x)  // Prints 2

	// 3. The new Function (Concept 2.3.3)
	// Creates a unnamed int, returns its address
	y := new(int)
	fmt.Println(*y) // Prints 0, (the "Zero Value")
	*y = 50
	fmt.Println(*y) // prints 50
}
