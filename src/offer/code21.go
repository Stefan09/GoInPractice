package offer

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/

func Exchange(nums []int) []int {
	if len(nums) < 1 || len(nums) > 50000 {
		return []int{}
	}
	//所有游标
	left, right := 0, len(nums)-1
	for left < right {
		for left < right {
			if nums[left]%2 != 0 {
				left++
			} else {
				break
			}
		}

		for left < right {
			if nums[right]%2 == 0 {
				right--
			} else {
				break
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
