package leetcode

/**
https://leetcode-cn.com/problems/two-sum/
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
*/

// brute force
func twoSum1(nums []int, target int) []int {
	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// upgrade
/*func twoSum2(nums []int, target int) []int {
	// sort
	sort.Ints(nums)

	// meet
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return nil
}
wrong: alter index
*/

// upgrade use additional space to save time
func twoSum3(nums []int, target int) []int {
	// build map to storage no pairs num
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i,j}
		}
		m[v] = i
	}

	return nil
}
