package main

import "fmt"

const (
	_ = iota
	A = iota
	B = 2
	C
)

func main() {
	fmt.Println(A, B, C)
}
