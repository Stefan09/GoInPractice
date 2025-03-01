package leetcode

import "sort"

/*字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]

提示：
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for _, str := range strs {
		chars := make([]byte, len(str))
		for i := 0; i < len(str); i++ {
			chars[i] = str[i]
		}
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		if _, ok := group[key]; ok {
			group[key] = append(group[key], str)
		} else {
			group[key] = []string{str}
		}
	}

	res := make([][]string, 0, len(group))
	for _, l := range group {
		res = append(res, l)
	}
	return res
}
