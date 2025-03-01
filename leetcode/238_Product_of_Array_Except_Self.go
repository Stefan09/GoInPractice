package leetcode

/*
	除自身以外数组的乘积

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:
输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：
2 <= nums.length <= 105
-30 <= nums[i] <= 30
输入 保证 数组 answer[i] 在  32 位 整数范围内

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
*/
func productExceptSelf(nums []int) []int {
	leftRes, rightRes := make([]int, len(nums)), make([]int, len(nums))
	leftRes[0], rightRes[len(nums)-1] = nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		leftRes[i] = leftRes[i-1] * nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightRes[i] = rightRes[i+1] * nums[i]
	}

	res := make([]int, len(nums))
	res[0], res[len(nums)-1] = rightRes[1], leftRes[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = leftRes[i-1] * rightRes[i+1]
	}
	return res
}