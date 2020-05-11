package offer

import "strings"

func ReplaceSpace(s string) string {
	strLength := len(s)
	if strLength == 0 || strLength >= 10000 {
		return ""
	}
	strResult := strings.ReplaceAll(s, " ", "%20")
	return strResult
}