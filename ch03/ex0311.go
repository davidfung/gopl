package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Ex 3.10")
	s := [...]string {"-1000000.9", "+100000.99", "10000.99", "1000.99", "100.99", "10.99", "1.9"}
	for _, v := range(s) {
		fmt.Printf("%15s\n", comma(v))
	}
}

func comma(s string) string {

	// optional sign
	sign := ""
	r := s
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = s[0:0]
		r = s[1:]
	}

	// floating point
    ss := strings.Split(r, ".")
	a := ss[0]
	b := ss[1]

	// thousand separator
	var buf bytes.Buffer
	var j int = 0
	for i := len(a) - 1; i >=0; i-- {
		j++
		buf.WriteByte(a[i])
		if j % 3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
	}

	t := buf.String()

	var rev bytes.Buffer
	for i := len(t)-1; i >= 0; i-- {
		rev.WriteByte(t[i])
	}

	return sign + rev.String() + "." + b

}