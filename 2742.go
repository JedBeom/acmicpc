package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 0
	fmt.Scan(&n)

	p := bufio.NewWriter(os.Stdout)

	for n > 0 {
		p.WriteString(strconv.Itoa(n) + "\n")
		n--
	}
	p.Flush()
}
