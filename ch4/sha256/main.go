// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("X"))
	diff := bitDiffCount(c1, c2)
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1, diff)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func bitDiffCount(s1, s2 [32]uint8) int {
	var count int
	for i := 0; i < len(s1); i++ {
		left := s1[i]
		right := s2[i]

		xor := left ^ right
		for xor != 0 {
			xor = xor & (xor - 1)
			count++
		}

	}

	return count
}

//!-
