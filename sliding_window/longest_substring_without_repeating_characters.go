package sliding_window

/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.

Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {
	var i, j, result int = 0, 0, 0
	charMap := make(map[uint8]int)

	for j < len(s) {
		if charIndex, found := charMap[s[j]]; found {
			i = charIndex + 1
		}
		charMap[s[j]] = j

		if (j - i) > result {
			result = j - i
		}

		j += 1
	}

	return result
}
