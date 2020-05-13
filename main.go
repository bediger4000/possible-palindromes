package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Create counts of each character appearing in the input string.
	counts := make(map[rune]int)
	for _, r := range []rune(os.Args[1]) {
		counts[r]++
	}
	// 2. Examine character counts, counting the number of odd character counts.
	oddCount := 0
	for _, count := range counts {
		oddCount += (count % 2)

		if oddCount > 1 {
			fmt.Printf("%q: false\n", os.Args[1])
		}
	}
	// 3. If at moust 1 count is odd, you've got a possible palindrome
	fmt.Printf("%q: true\n", os.Args[1])
}
