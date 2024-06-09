package main

import (
	"fmt"
)

var desc = "Write a version of rotate that operates in a single pass."

func main() {
	fmt.Println("Ex 4.4")
	fmt.Println(desc)

	fmt.Println("Rotate")
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	rotate(a)
	fmt.Println(a)

	fmt.Println("Rotate2")
	a = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	rotate2(a)
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}

func rotate2(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}
