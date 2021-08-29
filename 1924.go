package main

import (
	"fmt"
)

// 1924: 2007년.
// x y가 주어질 떄 x월 y일은 몇요일인지 출력하시오.
func main() {
	x, y := 0, 0
	fmt.Scanf("%d %d", &x, &y)
	months := []int{0, 3, 3, 6, 1, 4, 6, 2, 5, 0, 3, 5}
	weekdays := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

	today := (y + months[x-1]) % 7
	fmt.Println(weekdays[today])
}
