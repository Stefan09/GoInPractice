package leetcode

/**
https://leetcode-cn.com/problems/longest-palindromic-substring/
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
*/

/**
状态转移方程：s(i,j)=s(i+1,j-1)&&s[i]==s[j]
*/
func longestPalindrome(s string) string {
	l := len(s)
	if l == 1 || l == 2 && s[0] == s[1] {
		return s
	}

	// 初始化dp数组
	max := 1  // 当前最大长度
	left := 0 // 起始位置
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	// 计算dp数组
	// 错误原因：当前循环方式为从s头开始逐个子串判断是否回文，导致部分前驱状态没有计算
	// 正确方式：以长度划分最小子问题，j=l-1+i
	/*for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i == 1 { // 相邻字符
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 边计算边记录结果
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				left = i
			}
		}
	}
	fmt.Println(max, left, dp)*/

	for ll := 2; ll <= l; ll++ {
		for i := 0; i < l; i++ {
			j := ll - 1 + i
			if j >= l {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && ll > max {
				max = ll
				left = i
			}
		}
	}

	return s[left : max+left]
}
