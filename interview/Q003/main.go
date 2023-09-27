package q003

func reverseString(s string) (output string) {
	runes := []rune(s)

	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}
