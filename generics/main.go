package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[V Number](m map[string]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first": 12,
		"second": 34,
	}
	floats := map[string]float64{
		"first": 12.12,
		"secend": 34.34,
	}
	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}