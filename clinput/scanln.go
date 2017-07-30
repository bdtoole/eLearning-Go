package main

import "fmt"

func main() {

	word1, word2 := "", ""
	fmt.Print("Enter two words: ")
	count, err := fmt.Scanln(&word1, &word2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("You entered %v. There are %v words.\n", word1+" "+word2, count)
}
