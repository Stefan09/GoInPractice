package leetcode

/**
给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：
输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'
*/

func LongestValidParentheses(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	dp := make([]int, l)
	dp[0] = 0
	for i := 1; i < l; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i > dp[i-1] {
					if s[i-dp[i-1]-1] == '(' {
						if i >= dp[i-1]+2 {
							dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		} else {
			dp[i] = 0
		}
	}

	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}
