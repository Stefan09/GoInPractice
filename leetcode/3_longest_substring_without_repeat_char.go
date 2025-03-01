package leetcode

/*
	无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {
	left, right, maxSubLen := 0, 0, 0
	charSet := make(map[byte]int) // 窗口中字符统计 value=index

	for right < len(s) {
		char := s[right]
		if index, ok := charSet[char]; ok {
			maxSubLen = max(maxSubLen, right-left) // 先记结果
			for left < index {                     // 维护窗口中字符 为什么不是等号？字符重复出现不需要把重复字符消除掉
				delete(charSet, s[left])
				left++
			}
			left = index + 1      // 移动左指针越过重复字符 什么时候出现右边界和窗口内的元素重复
			charSet[char] = right // 维护下标；注意起初charSet是窗口中元素下标
		} else {
			maxSubLen = max(maxSubLen, right-left+1)
		}
		charSet[char] = right
		right++
	}
	return maxSubLen
}
