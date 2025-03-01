package leetcode

/* 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
// 连续非空 可使用前缀和
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	preSum := map[int]int{0: 1} // 子区间首个元素=k
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if c, ok := preSum[sum-k]; ok {
			count += c
		}
		preSum[sum]++ // 维护前缀和
	}
	return count
}
