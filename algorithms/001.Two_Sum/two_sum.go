package twoSum

func twoSum(nums []int, target int) []int {
	intMap := map[int]int{}

	for i, val1 := range nums {
		t := target - val1
		k, found := intMap[t]
		if found {
			return []int{k, i}
		} else {
			intMap[val1] = i
		}
	}
	return []int{}
}
