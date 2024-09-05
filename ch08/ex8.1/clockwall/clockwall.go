package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type City struct {
	name string
	url  string
	time string
}

func main() {
	cities := getCities()
	getCityTime(cities)
	showClockWall(cities)
}

func getCities() []City {
	var cities []City
	for _, x := range os.Args[1:] {
		y := strings.Split(x, "=")
		cities = append(cities, City{name: y[0], url: y[1]})
	}
	return cities
}

func getCityTime(cities []City) {
	for i := range cities {
		conn, err := net.Dial("tcp", cities[i].url)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		var b []byte

		_, err = conn.Read(b)
		if err != nil {
			cities[i].time = string(b)
		} else {
			cities[i].time = "n/a"
		}
	}
}

func showClockWall_old(cities []City) {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for _, city := range cities {
		fmt.Fprintf(tw, format, city.name, city.time.Format("15:04:05"))
	}
	tw.Flush()
}

func showClockWall(cities []City) {
	var clist []string
	var tlist []time.Time

	// 1st row: ciites
	for i := range cities {
		clist = append(clist, cities[i].name)
	}
	fmt.Printf("|")
	for _, city := range clist {
		fmt.Printf("%12s|", city)
	}
	fmt.Printf("\n")

	// 2nd row: time
	for i := range cities {
		tlist = append(tlist, cities[i].time)
	}
	fmt.Printf("|")
	for _, time := range tlist {
		fmt.Printf("%12s|", time.Format("15:04:05"))
	}
	fmt.Printf("\n")
}
