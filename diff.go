package main

import "fmt"
import "errors"

func lcsCheck(leftLimit int, leftFile []string, rightLimit int, rightFile []string) error {
	if leftLimit > len(leftFile) {
		message := fmt.Sprintf(
			"Out of bounds for left file: [%v] too large for slice of length [%v]",
			leftLimit,
			len(leftFile))
		return errors.New(message)
	}
	if rightLimit > len(rightFile) {
		message := fmt.Sprintf(
			"Out of bounds for right file: [%v] too large for slice of length [%v]",
			rightLimit,
			len(rightFile))

		return errors.New(message)
	}

	return nil
}

func lcs(leftLimit int, leftFile []string, rightLimit int, rightFile []string) (int, error) {
	if err := lcsCheck(leftLimit, leftFile, rightLimit, rightFile); err != nil {
		return 0, err
	}

	if leftLimit == 0 || rightLimit == 0 {
		return 0, nil
	}

	if leftFile[leftLimit-1] == rightFile[rightLimit-1] {
		lcs, err := lcs(leftLimit-1, leftFile, rightLimit-1, rightFile)
		if err != nil {
			return 0, err
		}

		return 1+lcs, nil
	}

	maxLeft, err := lcs(leftLimit-1, leftFile, rightLimit, rightFile)
	if err != nil {
		return 0, err
	}
	maxRight, err := lcs(leftLimit, leftFile, rightLimit-1, rightFile)
	if err != nil {
		return 0, err
	}

	if maxLeft > maxRight {
		return maxLeft, nil
	}
	return maxRight, nil
}

func lcsMatrix(leftFile []string, rightFile []string) ([][]int, error) {
	sequenceLengths := make([][]int, len(leftFile))

	for leftIndex, _ := range leftFile {
		sequenceLengths[leftIndex] = make([]int, len(rightFile))
		for rightIndex, _ := range rightFile {
			lr, err := lcs(leftIndex+1, leftFile, rightIndex+1, rightFile)
			sequenceLengths[leftIndex][rightIndex] = lr
			if err != nil {
				return [][]int{}, err
			}
		}
	}

	return sequenceLengths, nil
}

func main() {
	x, _ := lcs(0, []string{}, 0, []string{})
	fmt.Printf("%v ?= 0 - lcs(0, {}, 0, {})\n", x)

	x, _ = lcs(1, []string{"a"}, 1, []string{"a"})
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 1, {'a'})\n", x)

	x, _ = lcs(2, []string{"a", "b"}, 1, []string{"a"})
	fmt.Printf("%v ?= 1 - lcs(2, {'a', 'b'}, 1, {'a'})\n", x)

	x, _ = lcs(2, []string{"a", "b"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 2 - lcs(1, {'a', 'b'}, 2, {'a', 'b'})\n", x)

	x, _ = lcs(1, []string{"a"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 1 - lcs(1, {'a'}, 2, {'a', 'b'})\n", x)

	x, _ = lcs(1, []string{"b"}, 2, []string{"a", "b"})
	fmt.Printf("%v ?= 1 - lcs(1, {'b'}, 2, {'a', 'b'})\n", x)


	left := []string{"a", "b", "c", "d", "e", "f", "g"}
	right := []string{"w", "a", "b", "x", "y", "z", "e"}
	y, _ := lcsMatrix(right, left)
	fmt.Printf("  %v\n", left)
	for index, line := range y {
		fmt.Printf("%v %v\n", right[index], line)
	}
}


