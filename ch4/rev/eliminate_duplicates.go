// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//!+slice
	s := []string{"0", "1", "2", "3", "4", "4", "4", "5"}
	s = dedupe(s)
	fmt.Println(s) // "[0 1 2 3 4 5]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var fields []string
		for _, s := range strings.Fields(input.Text()) {
			fields = append(fields, s)
		}
		fields = dedupe(fields)
		fmt.Printf("%v\n", fields)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// remove adjacent duplicates from string slice.
func dedupe(s []string) []string {
	for i, j := 0, 1; j < len(s)-1; i, j = i+1, j+1 {
		if s[i] == s[j] {
			s = append(s[:i], s[j:]...)
			i--
			j--
		}
	}
	return s
}

//!-rev
