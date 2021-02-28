package offer

/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
f(n) = f(n-1)+f(n-2)
*/

//动态规划
func numWays(n int) int {
	if n < 0 || n > 100 {
		return -1
	}
	arr := make([]int, n+2)
	arr[0] = 0
	arr[1] = 1
	for iter := 2; iter <= n; iter++ {
		arr[iter] = (arr[iter-1] + arr[iter-2]) % (1e9 + 7)
	}
	return arr[n]
}
