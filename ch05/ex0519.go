package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

func foo() (x int) {
	defer func() {
		x = recover().(int)
	}()
	panic(6)
}
