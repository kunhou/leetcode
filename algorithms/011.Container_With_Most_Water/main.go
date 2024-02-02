package main

// https://leetcode.com/problems/container-with-most-water/
// 要解決這個問題，我們可以使用雙指針策略。具體方法是初始化兩個指針，
// 一個指向height數組的開始位置，另一個指向數組的結束位置。
// 每一步中，我們計算由這兩個指針所指向的線段和x軸形成的容器的面積，並記錄下找到的最大面積。
// 然後，我們移動指向較短線段的指針向另一個指針方向移動，因為移動指向較長線段的指針不會幫助我們找到更大的面積。
// 重複這個過程，直到兩個指針相遇。
func maxArea(height []int) int {
	var left, maxArea int
	right := len(height) - 1
	for left < right {
		width := right - left
		currentHigh := min(height[right], height[left])
		currentArea := width * currentHigh
		maxArea = max(maxArea, currentArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(s, t int) int {
	if s > t {
		return t
	}
	return s
}

func max(s, t int) int {
	if s > t {
		return s
	}
	return t
}
