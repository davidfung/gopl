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

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64 = 42
	return w, &count
}
func main() {
	var name = "Dolly"
	w, c := CountingWriter(os.Stdout)
	fmt.Fprintf(w, "hello, %s\n", name)
	fmt.Printf("Total bytes written = %d\n", *c)
}
