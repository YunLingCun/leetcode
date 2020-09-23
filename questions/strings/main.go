package main

import "fmt"

func contain(s string, c int32) int {
	for idx, char := range s {
		if c == char {
			return idx+1
		}
	}
	return 0
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var start, end = 0, 1
	var max = 1
	for _, c := range s[1:] {
		idx := contain(s[start:end], c)
		if idx == 0 {
			end += 1
			if end-start > max {
				max = end - start
			}
		} else {
			if end-start > max {
				max = end - start
			}
			start += idx
			end += 1
		}
	}
	return max
}

func main() {
	fmt.Printf("%d",lengthOfLongestSubstring("1111111111111111111111111122222233333444434565554433"))
}
