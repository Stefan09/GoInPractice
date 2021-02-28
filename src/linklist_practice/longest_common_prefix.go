package linklist_practice

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	commonPrefix := ""
	for i := 0; i < len(strs[0]); i++ {
		tempPrefix := strs[0][:i+1]
		for _, iter := range strs {
			if i > len(iter) || !strings.HasPrefix(iter, tempPrefix) {
				return commonPrefix
			}
		}
		commonPrefix = tempPrefix
	}

	return commonPrefix
}
