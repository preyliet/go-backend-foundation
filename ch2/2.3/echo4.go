package main

import (
	"flag"
	"fmt"
	"strings"
)

// n and sep are POINTERS (*bool and *string)
var n = flag.Bool("n", false, " omit trailing new line")
var sep = flag.String("s", " ", " separator")

func main() {
	flag.Parse() // Update variable froms from flags

	// We use *sep and *n to get the values the pointers point to
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
