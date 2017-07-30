package main

import "fmt"

func main() {

	a, b := 2, 2

	if a == b {
		fmt.Printf("%v and %v are equal\n", a, b)
	} else {
		fmt.Printf("%v and %v are not equal\n", a, b)
	}
}
