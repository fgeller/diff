package main

import "fmt"

// Let P_ij be the length of the longest subsequence common to the first i lines
// of the first file a and the first j lines of the second b.

func lcs(i int, a []string, j int, b []string) int {
	// TODO: err -- bounds check
	if i == 0 || j == 0 {
		return 0
	}

	if a[i-1] == b[j-1] {
		return 1 + lcs(i-1, a, j-1, b)
	}

	i1j := lcs(i-1, a, j, b)
	ij1 := lcs(i, a, j-1, b)
	if i1j > ij1 {
		return i1j
	}
	return ij1
}

func main() {
	fmt.Printf("%v ?= 0 - lcs(0, {}, 0, {})\n", lcs(0, []string{}, 0, []string{}))
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 1, {'a'})\n", lcs(1, []string{"a"}, 1, []string{"a"}))
	fmt.Printf("%v ?= 1 - lcs(2, {'a', 'b'}, 1, {'a'})\n", lcs(2, []string{"a", "b"}, 1, []string{"a"}))
	fmt.Printf("%v ?= 2 - lcs(1, {'a', 'b'}, 2, {'a', 'b'})\n", lcs(2, []string{"a", "b"}, 2, []string{"a", "b"}))
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 2, {'a', 'b'})\n", lcs(1, []string{"a"}, 2, []string{"a", "b"}))
	fmt.Printf("%v ?= 1 - lcs(1, {'b'}, 2, {'a', 'b'})\n", lcs(1, []string{"b"}, 2, []string{"a", "b"}))
}
