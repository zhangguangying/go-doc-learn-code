package main

import "fmt"

const (
	_ = iota
	A = iota
	B = 2
	C
)

func main() {
	m := map[string]int{}
	m["s"] = 1
	fmt.Println(m["m"])
}
