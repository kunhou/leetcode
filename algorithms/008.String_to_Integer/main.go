package main

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

// https://leetcode.com/problems/string-to-integer-atoi/description/
func myAtoi(s string) int {
	sign := 1
	result := 0

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
		result = result*10 + int(s[i]-'0')

		if result*sign > MaxInt32 {
			return MaxInt32
		}
		if result*sign < MinInt32 {
			return MinInt32
		}
		i++
	}
	result *= sign

	return result
}
