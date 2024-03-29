package leetcode

import "strings"

/**
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/
var numMap = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	res := make([]string, 0)

	if len(digits) == 0 {
		return res
	}
	if len(digits) == 1 {
		return strings.Split(numMap[digits[0]], "")
	}

	c := digits[0]
	s := LetterCombinations(digits[1:])
	for i := 0; i < len(numMap[c]); i++ {
		for _, iter := range s {
			res = append(res, string(numMap[c][i])+iter)
		}
	}
	return res
}
