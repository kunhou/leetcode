package main

func longestPalindrome(s string) string {
	return ""
}

// Start of Solution v1
func v1LongestPalindrome(s string) string {
	result := s[0:1]
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindromic(s[i:j+1]) && len(s[i:j+1]) > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func isPalindromic(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// End of Solution v1
