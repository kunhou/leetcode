package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {
	var a int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[a] {
			a++
			nums[a] = nums[i]
		}
	}
	return a + 1
}
