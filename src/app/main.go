package main

import (
	"fmt"

	"github.com/runnerdave/testing-example/src/minus"
	"github.com/runnerdave/testing-example/src/plus"
)

func main() {
	fmt.Println("Basic math operations")
	fmt.Printf("Sum 2 + 3: %d\n", plus.Plus(2, 3))
	fmt.Printf("Minus 3 - 2: %d\n", minus.Minus(3, 2))
}
