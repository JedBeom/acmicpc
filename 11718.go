package main

import (
	"fmt"
	"io"
)

func main() {
	n, i := "", ""
	for {
		_, err := fmt.Scan(&i)
		if err == io.EOF {
			break
		}
		n += i + "\n"
	}
	fmt.Println(n[:len(n)-2])
}
