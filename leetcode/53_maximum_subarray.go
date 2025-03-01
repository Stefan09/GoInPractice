package leetcode

import "math"

/*最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出:23


提示：
1 <= nums.length <= 10^5
-104 <= nums[i] <= 10^4

*/
// 不能暴力会超时
func maxSubArray(nums []int) int {
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

// 前缀和
func preSumMaxSubArray(nums []int) int {
	ans := math.MinInt
	minPreSum := 0
	sum := 0
	for _, x := range nums {
		sum += x                        // 当前的前缀和
		ans = max(ans, sum-minPreSum)   // 减去前缀和的最小值
		minPreSum = min(minPreSum, sum) // 维护前缀和的最小值
	}
	return ans
}
