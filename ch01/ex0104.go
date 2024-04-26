package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", names, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, names, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			s := ""
			for name := range names[line] {
				s += name + " "
			}
			fmt.Printf("%d:\"%s\" (%s)\n", n, line, strings.TrimSpace(s))
		}
	}
}

func countLines(f *os.File, name string, names map[string]map[string]int, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if names[t] == nil {
			names[t] = make(map[string]int)
		}
		names[t][name]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-