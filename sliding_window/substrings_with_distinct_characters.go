package sliding_window

/*
1876. Substrings of Size Three with Distinct Characters

A string is good if there are no repeated characters.
Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.
Note that if there are multiple occurrences of the same substring, every occurrence should be counted.
A substring is a contiguous sequence of characters in a string.

Example 1:
	Input: s = "xyzzaz"
	Output: 1
	Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz".
	The only good substring of length 3 is "xyz".

Example 2:
	Input: s = "aababcabc"
	Output: 4
	Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc", "bca", "cab", and "abc".
	The good substrings are "abc", "bca", "cab", and "abc".
*/

func CountGoodSubstrings(s string) int {
	result := 0

	if len(s) < 3 {
		return result
	}

	left := 0
	right := left + 3

	for right <= len(s) {
		currentStringGood := true
		currentWindowMap := make(map[uint8]int)

		for i := left; i < right; i++ {
			if _, found := currentWindowMap[s[i]]; found {
				currentStringGood = false

				break
			}
			currentWindowMap[s[i]] = i
		}

		if currentStringGood {
			result += 1
		}

		left += 1
		right += 1
	}

	return result
}
