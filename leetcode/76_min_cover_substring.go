package leetcode

/*
最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：
m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
*/
func MinWindow(s string, t string) string {
	/*
		窗口单向移动
		先扩张，移动right逐个字符放入窗口，更新计数器；直到覆盖字符数等于t的长度停止扩张
		再收缩，移动left逐个字符踢出窗口，更新计数器；直到覆盖字符数小于t的长度再次扩张
	*/
	start, length := 0, len(s)+1                              // 记录子串：起始位置 子串长度初始无限大
	left, right := 0, 0                                       // 窗口双端
	charSet, window := make(map[byte]int), make(map[byte]int) // 统计字符频次, 窗口中字符统计
	coverTimes := 0                                           // 覆盖字符数：字符和出现次数均匹配

	for i := 0; i < len(t); i++ {
		charSet[t[i]]++
	}

	for right < len(s) {
		if coverTimes < len(charSet) {
			window[s[right]]++
			if window[s[right]] == charSet[s[right]] { // 边维护边判断
				coverTimes++
			}
		}
		right++
		for coverTimes == len(charSet) {
			if right-left < length {
				start, length = left, right-left
			}
			if window[s[left]] == charSet[s[left]] {
				coverTimes--
			}
			window[s[left]]--
			left++
		}
	}
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
