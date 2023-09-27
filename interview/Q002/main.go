package main

func isUniqString(str string) bool {
	if len(str) == 0 || len(str) > 3000 {
		return false
	}
	var mark1, mark2, mark3, mark4 uint64
	for _, r := range str {
		n := uint64(r)
		var mark *uint64
		switch n / uint64(64) {
		case 0:
			mark = &mark1
		case 1:
			mark = &mark2
		case 2:
			mark = &mark3
		case 3:
			mark = &mark4
		}
		n = n % 64

		if (*mark)&(1<<n) != 0 {
			return false
		}
		*mark = *mark | uint64(1<<n)
	}
	return true
}
