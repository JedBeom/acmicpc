package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		fmt.Print(strings.Repeat(" ", n-i), strings.Repeat("*", i), "\n")
	}
}
