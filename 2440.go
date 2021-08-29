package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	for n > 0 {
		fmt.Println(strings.Repeat("*", n))
		n -= 1
	}
}
