// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fmt.Println(s)
	// Rotate s left by two positions.
	s = rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, shift int) []int {
	if shift < 0 || len(s) == 0 {
		return s
	}

	r := len(s) - shift%len(s)
	s = append(s[r:], s[:r]...)
	return s

	//the following does not handle both odd & even len(s)
	//for i, v := range s[:shift] {
	//	for j := i + len(s) - shift; j >= 0; j -= shift {
	//		fmt.Printf("i=%d, j=%d, v=%d, s=%v\n", i, j, v, s)
	//		s[j], v = v, s[j]
	//		fmt.Printf("i=%d, j=%d, v=%d, s=%v\n", i, j, v, s)
	//	}
	//}
}

//!-rev
