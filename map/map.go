package main

import "fmt"

func main() {

	/*
		a map is a table of key/value pairs. maprs are
		also known as hash tables, dictonaries, or
		associative arrays in other languages.

		the keys of a map must be "hashable," meaning
		that a unique value can be assigned to each key
		and that equality can be evaluated for every key.

		ints, floats, strings, etc. are all hashable types,
		but arrays, slices, etc. are not.

		the values of a map need not be hashable.
	*/

	digits := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	digits["five"] = 5

	if digits["four"] != 0 {
		fmt.Printf("that digit is %v\n", digits["four"])
	} else {
		fmt.Println("that digit is not in the digits map.")
	}

	// this is called the comma ok idiom. It test for
	// the presence of a pair in the map

	key := "six"

	var digit int
	var ok bool

	if digit, ok = digits[key]; ok {
		fmt.Printf("%v is %v\n", key, digit)
	} else {
		fmt.Printf("%v is not a key in the map.\n", key)
	}

	// we can also throw away the value if we're only
	// interested in knowing if they key exists.
	// this is done to avoid the problem of initializing
	// but not using a variable.

	if _, ok = digits["nine"]; ok {
		fmt.Println("nine is in the map.")
	} else {
		fmt.Println("nine is not in the map.")
	}

	// deleting a pair from a map:

	delete(digits, "zero")

	// we can print the entire map:
	fmt.Printf("%v\n", digits)

}
