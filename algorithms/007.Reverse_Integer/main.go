package main

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go
// outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > MaxInt32 || res < MinInt32 {
		return 0
	}
	return res
}
