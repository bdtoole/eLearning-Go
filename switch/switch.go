package main

import "fmt"

func main() {

	grades := [7]int{93, 72, 68, 54, 100, 43, 86}
	strScore := ""
	for score := range grades {
		switch {
		case grades[score] > 90:
			strScore = "A"
		case grades[score] > 80:
			strScore = "B"
		case grades[score] > 70:
			strScore = "C"
		case grades[score] > 60:
			strScore = "D"
		default:
			strScore = "F"
		}
		fmt.Printf("%v, letter grade %v\n", grades[score], strScore)
	}
}
