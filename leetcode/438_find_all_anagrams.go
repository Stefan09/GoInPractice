package leetcode

/*
	找到所有字母异位词

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/
// 当len(s)==len(p)时，退化成O(n^2)
func findAnagrams(s string, p string) []int {
	var res []int
	left, right := 0, len(p)
	charSet, window := make(map[uint8]int), make(map[uint8]int)

	for i := 0; i < len(p); i++ {
		charSet[p[i]]++
	}

	for right <= len(s) {
		ss := s[left:right]
	L1:
		for i := 0; i < len(ss); i++ {
			if _, ok := charSet[ss[i]]; !ok {
				break
			} else {
				window[ss[i]]++
			}
			if i == len(ss)-1 {
				for char, num := range charSet {
					if window[char] != num {
						break L1
					}
				}
				res = append(res, left)
			}
		}
		left++
		right++
		clear(window)
	}
	return res
}
