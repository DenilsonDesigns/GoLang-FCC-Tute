package main

import (
	"fmt"
)

var (
	i int = 42
	j int = 64
	k int = 72
)

func main() {
	fmt.Printf("%v, %T", i, i)
	// fmt.Printf("%v, %T", j, j)
	// fmt.Printf("%v, %T", k, k)
}
