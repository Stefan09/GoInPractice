package leetcode

import "fmt"

/**
W 背包承重
N 物品数量
wt 物品重量
val 物品价值
*/
func Knapsack(W, N int, wt, val []int) {
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	fmt.Println(dp)

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j-wt[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
		}
	}
	fmt.Println(dp[N][W])
}
