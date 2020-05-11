package offer

import "sort"

//利用map
func FindRepeatNumber1(nums []int) int {
	counterMap := make(map[int]int)
	for _, iter := range nums {
		if _, ok := counterMap[iter]; ok {
			return iter
		} else {
			counterMap[iter]++
		}
	}

	return 0
}

//sort
func FindRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for index, value := range nums {
		if nums[index] == nums[index+1] {
			return value
		}
	}
	return 0
}