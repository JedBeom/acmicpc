package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Print(strings.Repeat(" ", i), strings.Repeat("*", n-i), "\n")
	}
}
