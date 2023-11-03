package string_practice

import "fmt"

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例 1：

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
示例 2：

输入: s1= "ab" s2 = "eidboaoo"
输出: False


提示：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	//needle字符集
	needleCharSet, substrCharSet := [26]int{}, [26]int{}
	//滑动窗口
	left, right := 0, len(s1)-1

	//整理needle字符集
	for iter := 0; iter < len(s1); iter++ {
		needleCharSet[s1[iter]-'a']++
		substrCharSet[s2[iter]-'a']++
	}

	//使用滑动窗口
	for right < len(s2) {
		if match(needleCharSet[:], substrCharSet[:]) {
			return true
		}

		right++
		if right >= len(s2) {
			return false
		}
		substrCharSet[s2[right]-'a']++

		substrCharSet[s2[left]-'a']--
		left++
	}

	return false
}

func match(set1, set2 []int) bool {
	if len(set1) != len(set2) {
		return false
	}

	for i := 0; i < len(set1); i++ {
		if set1[i] != set2[i] {
			return false
		}
	}

	return true
}

//递归全排列
func RecursionArrange(arr []int, cur int) {
	if len(arr) == 0 {
		return
	}

	if cur == len(arr)-1 {
		fmt.Println(arr)
		return
	}

	for iter := cur; iter < len(arr); iter++ {
		arr[cur], arr[iter] = arr[iter], arr[cur]
		RecursionArrange(arr, cur+1)
		arr[cur], arr[iter] = arr[iter], arr[cur]
	}

}
