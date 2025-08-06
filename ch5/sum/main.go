// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"errors"
	"fmt"
)

// !+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//!-

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("no values found")
	}
	var m int
	for _, val := range vals {
		if m < val {
			m = val
		}
	}
	return m, nil
}

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	m, err := max(values...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	//!-slice
}
