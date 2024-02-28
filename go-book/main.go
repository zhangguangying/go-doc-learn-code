package main

import "fmt"

const (
	_ = iota
	A = iota
	B = 2
	C
)

func main() {
	l1 := [256]int{'a': 1, 'b': 2}
	fmt.Println(l1)
}
