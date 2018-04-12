package logestSubstring

func lengthOfLongestSubstring(s string) int {
	res := 0
	leftBound := 0
	dic := map[rune]int{}

	for i, ch := range s {
		p, found := dic[ch]
		if found && p >= leftBound {
			leftBound = p + 1
		}
		dic[ch] = i
		count := i - leftBound + 1
		if res < count {
			res = count
		}
	}
	return res
}
