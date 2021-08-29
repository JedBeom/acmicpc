package main

import (
	"fmt"
	"strconv"
)

func st(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func main() {
	n, a := 0, ""
	fmt.Scan(&n)
	fmt.Scan(&a)
	y := 0
	for _, value := range a {
		y += st(string(value))
	}
	fmt.Println(y)
}
