package sliding_window

import "math"

/*
Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:
	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Example 3:
	Input: s = "a", t = "aa"
	Output: ""
	Explanation: Both 'a's from t must be included in the window.
	Since the largest window of s only has one 'a', return empty string.
*/

func MinWindow(s string, t string) string {
	result := ""

	if len(s) == 0 || len(t) == 0 {
		return result
	}

	charFrequencies := make(map[uint8]int)
	for _, char := range t {
		charFrequencies[uint8(char)] += 1
	}

	lettersToSatisfy := len(charFrequencies)
	left := 0
	minLeft, minSize := 0, math.MaxInt32

	for j := 0; j < len(s); j++ {
		char := s[j]
		if _, found := charFrequencies[char]; found {
			charFrequencies[char] -= 1

			if charFrequencies[char] == 0 {
				lettersToSatisfy -= 1
			}
		}

		for lettersToSatisfy == 0 {
			if j-left+1 < minSize {
				minSize = j - left + 1
				minLeft = left
			}

			leftChar := s[left]
			if _, found := charFrequencies[leftChar]; found {
				charFrequencies[leftChar] += 1

				if charFrequencies[leftChar] > 0 {
					lettersToSatisfy += 1
				}
			}

			left += 1
		}
	}

	if minSize == math.MaxInt32 {
		return result
	}

	return string([]rune(s)[minLeft : minLeft+minSize])
}
