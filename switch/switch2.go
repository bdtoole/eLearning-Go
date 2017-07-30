package main

import "fmt"

func main() {

	numSize := ""

	//loop from 1 to 10:
	for i := 1; i <= 10; i++ {
		//assign a numSize to i
		if i < 4 {
			numSize = "small"
		} else if i < 8 {
			numSize = "medium"
		} else {
			numSize = "large"
		}

		//switch on the numSize
		switch numSize {
		case "small":
			fmt.Printf("%v is a small number.\n", i)
		case "medium":
			fmt.Printf("%v is bigger, but not large.\n", i)
		default:
			fmt.Printf("%v is large.\n", i)
		}
	}
}
