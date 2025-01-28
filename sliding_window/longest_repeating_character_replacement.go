package sliding_window

/*
You are given a string s and an integer k.
You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
	Input: s = "ABAB", k = 2
	Output: 4
	Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
	Input: s = "AABABBA", k = 1
	Output: 4
	Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
	The substring "BBBB" has the longest repeating letters, which is 4.
	There may exists other ways to achieve this answer too.
*/

func CharacterReplacement(s string, k int) int {
	maxLength := 0
	maxFrequency := 0
	currentLength := 0
	currentFrequency := 0
	frequencies := make(map[uint8]int)
	i := 0

	for j := 0; j < len(s); j++ {
		currentFrequency = findCharacterFrequency(s[j], frequencies)
		frequencies[s[j]] = currentFrequency

		if currentFrequency > maxFrequency {
			maxFrequency = currentFrequency
		}

		currentLength += 1

		if currentLength-maxFrequency > k {
			if frequency, found := frequencies[s[i]]; found {
				frequencies[s[i]] = frequency - 1
			}
			i += 1
			currentLength -= 1

			continue
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func findCharacterFrequency(char uint8, frequencies map[uint8]int) int {
	if frequency, found := frequencies[char]; found {
		return frequency + 1
	} else {
		return 1
	}
}
