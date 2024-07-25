// The strings.NewReader function returns a value that satisifies the io.Reader interface
// (and others) by reading from its argument, a string.  Implement a simple version of
// NewReader yourself, and use it to make the HTML parser($5.2) take input from a string.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type MyString struct {
	s   string
	ptr int
}

func (myString *MyString) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(myString.s[myString.ptr:]))
	myString.ptr += n
	if myString.ptr >= len(myString.s) {
		err = io.EOF
	}
	fmt.Printf("s=%d p=%d n=%d\n", len(myString.s), len(p), n)
	return n, err
}

func (myString *MyString) newReader(s string) io.Reader {
	myString.s = s
	return myString
}

// !+
func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	s = s + s + s + s
	s = s + s + s + s
	s = s + s + s + s
	myString := MyString{s: s}
	doc, err := html.Parse(myString.newReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
