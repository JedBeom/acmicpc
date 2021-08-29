package main

import (
	"fmt"
	"os"
)

func main() {
	n := 0
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		fmt.Fprintf(os.Stdout, "%d * %d = %d\n", n, i, n*i)
	}
}
