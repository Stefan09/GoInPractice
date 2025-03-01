package leetcode

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
示例 1：
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

示例 2：
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。

提示：
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]bool, sum+1)
		for j := 0; j < sum+1; j++ {
			if j == 0 {
				dp[i][j] = true
			}
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < sum+1; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
		}
	}
	return dp[len(nums)][sum]
}
