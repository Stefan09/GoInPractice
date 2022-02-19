package leetcode

import (
	"container/list"
	"fmt"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例2：
输入：s = "()[]{}"
输出：true

示例3：
输入：s = "(]"
输出：false

示例4：
输入：s = "([)]"
输出：false

示例5：
输入：s = "{[]}"
输出：true

提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := list.New()
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack.PushFront(s[i])
		} else {
			if e := stack.Front(); e == nil {
				return false
			} else {
				l := e.Value.(uint8)
				if match(l, s[i]) {
					stack.Remove(stack.Front())
				} else {
					return false
				}
			}
		}
	}

	return stack.Len() == 0
}

func isLeft(c uint8) bool {
	return c == '(' || c == '[' || c == '{'
}

func isRight(c uint8) bool {
	return c == ')' || c == ']' || c == '}'
}

func match(l, r uint8) bool {
	switch r {
	case ')':
		return l == '('
	case ']':
		return l == '['
	case '}':
		return l == '{'
	default:
		return false
	}
}

func visit(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
