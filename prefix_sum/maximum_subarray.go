package prefixsum

import "math"

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:
	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
	Input: nums = [1]
	Output: 1
	Explanation: The subarray [1] has the largest sum 1.

Example 3:
	Input: nums = [5,4,-1,7,8]
	Output: 23
	Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:
	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
*/

func MaxSubArray(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]

	for _, num := range nums[1:] {
		currentSum = int(math.Max(float64(currentSum+num), float64(num)))

		if currentSum > maxSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}
