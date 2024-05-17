package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Ex 3.10")
	s := [...]string {"1000000", "100000", "10000", "1000", "100", "10", "1"}
	for _, v := range(s) {
		fmt.Printf("%10s\n", comma(v))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	var j int = 0
	for i := len(s) - 1; i >=0; i-- {
		j++
		buf.WriteByte(s[i])
		if j % 3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
	}

	t := buf.String()

	var rev bytes.Buffer
	for i := len(t)-1; i >= 0; i-- {
		rev.WriteByte(t[i])
	}

	return rev.String()

}