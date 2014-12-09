package main

import "fmt"
import "errors"

func longestCommonSubsequence(leftLimit int, leftFile []string, rightLimit int, rightFile []string) (int, error) {
	if leftLimit == 0 || rightLimit == 0 {
		return 0, nil
	}
	if leftLimit > len(leftFile) {
		message := fmt.Sprintf(
			"Out of bounds for left file: [%v] too large for slice of length [%v]",
			leftLimit,
			len(leftFile)
		)
		return 0, errors.New(message)
	}
	if rightLimit > len(rightFile) {
		message := fmt.Sprintf(
			"Out of bounds for right file: [%v] too large for slice of length [%v]",
			rightLimit,
			len(rightFile)
		)
		return 0, errors.New(message)
	}



	if leftFile[leftLimit-1] == rightFile[rightLimit-1] {
		lcs, err := longestCommonSubsequence(leftLimit-1, leftFile, rightLimit-1, rightFile)
		if err != nil {
			return 0, err
		}

		return 1+lcs, nil
	}

	maxLeft, err := longestCommonSubsequence(leftLimit-1, leftFile, rightLimit, rightFile)
	if err != nil {
		return 0, err
	}
	maxRight, err := longestCommonSubsequence(leftLimit, leftFile, rightLimit-1, rightFile)
	if err != nil {
		return 0, err
	}

	if maxLeft > maxRight {
		return maxLeft, nil
	}
	return maxRight, nil
}

func main() {
	lcs, _ := longestCommonSubsequence(0, []string{}, 0, []string{})
	fmt.Printf("%v ?= 0 - lcs(0, {}, 0, {})\n", lcs)

	lcs, _ = longestCommonSubsequence(1, []string{"a"}, 1, []string{"a"})
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 1, {'a'})\n", lcs)

	lcs, _ = longestCommonSubsequence(2, []string{"a", "b"}, 1, []string{"a"})
	fmt.Printf("%v ?= 1 - lcs(2, {'a', 'b'}, 1, {'a'})\n", lcs)

	lcs, _ = longestCommonSubsequence(2, []string{"a", "b"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 2 - lcs(1, {'a', 'b'}, 2, {'a', 'b'})\n", lcs)

	lcs, _ = longestCommonSubsequence(1, []string{"a"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 2, {'a', 'b'})\n", lcs)

	lcs, _ = longestCommonSubsequence(1, []string{"b"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 1 - lcs(1, {'b'}, 2, {'a', 'b'})\n", lcs)
}
