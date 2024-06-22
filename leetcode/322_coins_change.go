package leetcode

/*
*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0

提示：
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4
*/
// 超出时间限制
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := amount + 1
	for _, coin := range coins {
		tmp := CoinChange(coins, amount-coin)
		if tmp == -1 {
			continue
		}
		res = min(res, tmp+1)
	}
	if res == amount+1 {
		return -1
	}
	return res
}

// 优化自上而下的执行效率 使用local cache缓存结果避免重复计算
func dp(coins []int, amount int, lc map[int]int) int {
	if val, ok := lc[amount]; ok {
		return val
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := amount + 1
	for _, coin := range coins {
		tmp := dp(coins, amount-coin, lc)
		if tmp == -1 {
			continue
		}
		res = min(res, tmp+1)
	}
	if res == amount+1 {
		res = -1
	}
	lc[amount] = res
	return res
}

func CoinChangeV2(coins []int, amount int) int {
	lc := make(map[int]int)
	return dp(coins, amount, lc)
}

// 动态规划 自下而上
func CoinChangeV3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= len(dp)-1; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
