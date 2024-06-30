package leetcode

/*
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false

提示：
1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
*/

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	left, right := 0, 0                                      // 左右指针
	matchTimes := 0                                          // 匹配计数
	target, window := make(map[byte]int), make(map[byte]int) // 子串字符计数 窗口计数

	for i := 0; i < len(s1); i++ {
		target[s1[i]]++
	}
	for right < len(s2) {
		if _, ok := target[s2[right]]; ok {
			window[s2[right]]++
			if window[s2[right]] == target[s2[right]] {
				matchTimes++
			}
		}
		right++
		for matchTimes == len(target) {
			if right-left == len(s1) {
				return true
			}
			if _, ok := target[s2[left]]; ok {
				if window[s2[left]] == target[s2[left]] {
					matchTimes--
				}
				window[s2[left]]--
			}
			left++
		}
	}
	return false
}
