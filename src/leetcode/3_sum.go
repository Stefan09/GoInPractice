package leetcode

import "sort"

/**
https://leetcode-cn.com/problems/3sum/
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]

提示：
0 <= nums.length <= 3000
-105 <= nums[i] <= 105
 */

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}

	// brute force time out because of O(n^3)
	// use double pointer
	sort.Ints(nums)
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1,len(nums)-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 结果集去重 i,left,right不可遗漏
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return res
}
