package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// Refer to Method 3
// SEE: https://www.geeksforgeeks.org/program-for-nth-fibonacci-number/
func fibonacci() func() int {
	left := -1
	right := 1
	return func() int {
		current := left + right
		left = right
		right = current
		return current
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}