package two_pointers

import "strings"

/*
https://leetcode.com/problems/valid-palindrome/description/

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
	Input: s = "A man, a plan, a canal: Panama"
	Output: true
	Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
	Input: s = "race a car"
	Output: false
	Explanation: "raceacar" is not a palindrome.

Example 3:
	Input: s = " "
	Output: true
	Explanation: s is an empty string "" after removing non-alphanumeric characters.
	Since an empty string reads the same forward and backward, it is a palindrome.
*/

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	cleaned := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			cleaned += string(char)
		}
	}

	left, right := 0, len(cleaned)-1
	isPallindrome := true

	for left <= right {
		if cleaned[left] != cleaned[right] {
			isPallindrome = false
		}

		left++
		right--
	}

	return isPallindrome
}
