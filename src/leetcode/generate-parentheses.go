package leetcode

/**
https://leetcode-cn.com/problems/generate-parentheses/

数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/

func generateParenthesis(n int) []string {
	res := new([]string)
	backtrack(n, n, "", res)
	return *res
}

func backtrack(l, r int, s string, res *[]string) {
	if l > 0 {
		backtrack(l-1, r, s+"(", res)
	}
	if r > l {
		s += ")"
		r--
		if r == l && r == 0 {
			*res = append(*res, s)
			return
		}
		backtrack(l, r, s, res)
	}
}
