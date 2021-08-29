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

	p := bufio.NewWriterSize(os.Stdout, 100000)

	i := 1
	for i <= n {
		p.WriteString(strconv.Itoa(i) + "\n")
		i++
	}
	p.Flush()
}
