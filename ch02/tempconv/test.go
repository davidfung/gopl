package main

import "fmt"
import "amg99.com/tempconv"

func main() {
	fmt.Println("Hello")
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC);
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}