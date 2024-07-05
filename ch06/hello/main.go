package main

import "fmt"

type Integer int

func (i Integer) Hello() {
	fmt.Printf("Hello, I am integer %d\n", i)
}

func main() {
	var i Integer
	i = 42
	fmt.Printf("int %d\n", i)
	i.Hello()
	(i + 1).Hello()
}
