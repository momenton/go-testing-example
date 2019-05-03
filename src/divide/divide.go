package divide

import "fmt"

func Divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem with Divide function - ", r)
		}
	}()
	return a / b
}
