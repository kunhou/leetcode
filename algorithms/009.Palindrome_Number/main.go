package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNum int
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x = x / 10
	}
	return x == revertedNum || x == revertedNum/10
}
