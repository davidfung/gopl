package main

import "fmt"

const (
	KB int = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	// ZB = EB * 1000
	// YB = ZB * 1000
)

func main() {
	fmt.Printf("%T %#[1]v\n", KB)
	fmt.Printf("%T %#[1]v\n", MB)
	fmt.Printf("%T %#[1]v\n", GB)
	fmt.Printf("%T %#[1]v\n", PB)
	fmt.Printf("%T %#[1]v\n", EB)
	// fmt.Printf("%T %#[1]v\n", ZB)
	// fmt.Printf("%T %#[1]v\n", YB)
}