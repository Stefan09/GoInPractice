package leetcode

/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:
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
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

示例 4:
输入: s = ""
输出: 0

提示：
0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/

/**
 * 1.滑动窗口，跳至重复位置滑动
 * 2.map用作set
 * 3.循环有多个出口时需要判断
 * 4.字符串中遍历字符，通过index获得1byte，通过range获得rune(int32)
 */
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	max := 1
	for i := 0; i < len(s); {
		set := make(map[uint8]int)
		set[s[i]] = i

		j := i + 1
		for ; j < len(s); j++ {
			if index, ok := set[s[j]]; ok {
				i = index + 1
				break
			}
			set[s[j]] = j
		}
		if len(set) > max {
			max = len(set)
		}

		// 双出口判断（包含了内循环条件不成立的情况）
		if j >= len(s) {
			break
		}
	}

	return max
}
