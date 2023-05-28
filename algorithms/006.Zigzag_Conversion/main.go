package main

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
// of rows like this: (you may want to display this pattern in a fixed font
// for better legibility)
//
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// P   A   H   N
// A P L S I I G
// Y   I   R

func convert(s string, numRows int) string {
	var result string
	if numRows <= 1 {
		return s
	}
	rows := make([]string, numRows)

	r := 0 // current row
	nextDirection := -1

	for i := 0; i < len(s); i++ {
		rows[r] += string(s[i])
		if r == 0 || r == numRows-1 {
			nextDirection *= -1
		}
		r += nextDirection
	}

	for i := 0; i < numRows; i++ {
		result += rows[i]
	}

	return result
}
