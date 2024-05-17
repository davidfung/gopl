package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "abcdefg"
	s2 := "cdefgab"
	b := isAnagram(s1, s2)
	fmt.Printf("s1=%s, s2=%s, %v\n", s1, s2, b)
}

func isAnagram(s1 string, s2 string) bool {
    s3 := sortString(s1)
    s4 := sortString(s2)
	return s3 == s4
}

func sortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	u := strings.Join(t, "")
	return u
}