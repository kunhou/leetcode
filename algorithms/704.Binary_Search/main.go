package main

// https://leetcode.com/problems/binary-search/
// Given an array of integers nums which is sorted in ascending order,
// and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

func search(nums []int, target int) int {
	for i, val := range nums {
		if val == target {
			return i
		}
	}
	return -1
}
