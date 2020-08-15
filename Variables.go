package main

import (
	"fmt"
	"strconv"
)

var (
	i int = 42
	j int = 64
	k int = 72
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
