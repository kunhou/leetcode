package main

// Given an integer array nums, find the contiguous subarray
//  (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// Example 2:
// Input: nums = [1]
// Output: 1

// Example 3:
// Input: nums = [5,4,-1,7,8]
// Output: 23

// maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];
func maxSubArray(nums []int) int {
	d := make([]int, len(nums))
	d[0] = nums[0]
	max := d[0]

	for i := 1; i < len(nums); i++ {
		sum := nums[i]
		if d[i-1] > 0 {
			sum += d[i-1]
		}
		d[i] = sum
		if sum > max {
			max = d[i]
		}
	}

	return max
}
