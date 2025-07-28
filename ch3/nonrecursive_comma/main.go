// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", formatNumericString(os.Args[i]))
	}
}

func formatNumericString(s string) string {
	var buf bytes.Buffer
	signed := strings.Contains(s[:1], "+") || strings.Contains(s[:1], "-")
	afterSign := s
	if signed {
		buf.WriteString(s[:1])
		afterSign = s[1:]
	}
	point := strings.LastIndex(afterSign, ".")
	var afterPoint string
	beforePoint := afterSign
	if point != -1 {
		beforePoint = afterSign[:point]
		afterPoint = afterSign[point:]
	}
	buf = comma(beforePoint, buf)
	buf.WriteString(afterPoint)
	return buf.String()
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string, buf bytes.Buffer) bytes.Buffer {
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return buf
	}
	var remainder = n % 3
	if remainder > 0 {
		buf.WriteString(s[:remainder])
		buf.WriteString(",")
	}
	for i := remainder; i < n; i = i + 3 {
		buf.WriteString(s[i : i+3])
		if i < n-3 {
			buf.WriteString(",")
		}
	}
	return buf
}

//!-
