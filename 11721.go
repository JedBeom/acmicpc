package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	n := 10
	for n < len(s) {
		fmt.Println(s[:n])
		s = s[n:]
	}
	if len(s) > 0 {
		fmt.Println(s)
	}
}
