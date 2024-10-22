package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialise a map for integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialise a map for float values
	floats := map[string]float64{
		"first":  44.44,
		"second": 33.33,
	}

	fmt.Printf("Non generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	fmt.Printf("Generic sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	fmt.Printf("Generic sums with constraints: %v and %v\n",
		sumNumbers(ints),
		sumNumbers(floats),
	)
}

// Sums values of map, supporting both integers and floats
func sumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Adds together the value of m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Support int64 and float64 types
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
