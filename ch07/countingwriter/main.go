// Ex 7.2 CountingWriter
// See page 174.

// Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 value
// that at any moment contains the number of bytes written to the new Writer.

package main

import (
	"fmt"
	"io"
	"os"
)

type MyWriter struct {
	Count      int64
	OrigWriter io.Writer
}

func (c *MyWriter) Write(p []byte) (n int, err error) {
	n, err = c.OrigWriter.Write(p)
	c.Count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	myWriter := MyWriter{OrigWriter: w}
	return &myWriter, &myWriter.Count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprintln(writer, "hello, world")
	fmt.Printf("Total bytes written = %d\n", *count)
	fmt.Fprintf(writer, "%s\n", "Safety")
	fmt.Printf("Total bytes written = %d\n", *count)
}
