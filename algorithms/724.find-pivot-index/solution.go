package main

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, val := range nums {
		sum += val
	}
	for i, val := range nums {
		rightSum := sum - leftSum - val
		if leftSum == rightSum {
			return i
		}
		leftSum += val
	}
	return -1
}
