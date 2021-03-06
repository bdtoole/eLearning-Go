package main

import "fmt"

func main() {
	fmt.Printf("5! = %v\n", fact(5))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}