package main

import "fmt"

func main() {
	a := 2
	b := 3
	s := Sum(a, b)
	fmt.Println(s)
}
func Sum(first, second int) int {
	Sum := first + second
	return Sum
}
