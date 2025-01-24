package sliding_window

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than (10^-5) will be accepted.


Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:
Input: nums = [5], k = 1
Output: 5.00000
*/

func FindMaxAverage(nums []int, k int) float64 {
	var result, sum = 0, 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	result = sum

	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > result {
			result = sum
		}
	}

	return float64(result) / float64(k)
}
