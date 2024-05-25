// Ex 14.1
package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	fmt.Println("Ex 4.1")
	// sha256 hash
	//                        1               2               3
	//        0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
	// "aaa" hash1 := "9834876dcfb05cb167a5c24953eba58c4ac89b1adf57f28f2f9d09af107ee8f0"
	// "bbb" hash2 := "3cf9a1a81f6bdeaf08a343c1e1c73e89cf44c06ac2427a892382cae825e7c9c1"

	hash1 := sha256.Sum256([]byte("aaa"))
	hash2 := sha256.Sum256([]byte("bbb"))
	fmt.Printf("hash1 (%T)=%x\n", hash1, hash1)
	fmt.Printf("hash2 (%T)=%x\n", hash2, hash2)

	cnt := countDiffBits(hash1, hash2)
	fmt.Printf("hash1 = %x\nhash2 = %x\ndiff bits = %d\n", hash1, hash2, cnt)
}

func countDiffBits(hash1, hash2 [32]uint8) int {
	//todo
	h1 := [][]uint8{hash1[:8], hash1[8:16], hash1[16:24], hash1[24:32]}
	h2 := [][]uint8{hash2[:8], hash2[8:16], hash2[16:24], hash2[24:32]}
	count := 0
	for i := range h1 {
		fmt.Println(i)
		fmt.Printf("h1[%d], %v\n", i, h1[i])
		fmt.Printf("h2[%d], %v\n", i, h2[i])
		fmt.Printf("%T\n", h1[i])

		var u1 uint64
		for _, b := range h1[i] {
			u1 = u1 << 8 + uint64(b)
		}
		var u2 uint64
		for _, b := range h2[i] {
			u2 = u2 << 8 + uint64(b)
		}
	    u := u1 ^ u2
	    count += PopCount(u)
	}
	return count
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}