package leetcode

/**
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
进阶：
你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109
*/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)

	if n == 0 {
		return res
	}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			for l := mid; l >= 0 && nums[l] == nums[mid]; l-- {
				res[0] = l
			}
			for r := mid; r <= n-1 && nums[r] == nums[mid]; r++ {
				res[1] = r
			}
			break
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
