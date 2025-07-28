package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	anagram := isAnagram(os.Args[1], os.Args[2])
	fmt.Printf("  %t\n", anagram)
}

func isAnagram(s string, s2 string) bool {
	if len(s) != len(s2) {
		return false
	}
	r1 := []rune(s)
	r2 := []rune(s2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	return string(r1) == string(r2)
}
