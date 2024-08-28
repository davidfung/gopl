package main

import (
	"fmt"
	"os"
	"strings"
)

type City struct {
	name string
	url  string
}

func main() {
	fmt.Println("Hello clockwall")
	var cities []City
	for _, x := range os.Args[1:] {
		y := strings.Split(x, "=")
		cities = append(cities, City{name: y[0], url: y[1]})
	}
	showClockWall(cities)
}

func showClockWall(cities []City) {
	fmt.Printf("%v\n", cities)
}
