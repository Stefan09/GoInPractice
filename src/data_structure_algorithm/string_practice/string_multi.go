package string_practice

/**
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

//思路：使用乘法分配律
func Multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//两个字符串按位拆解
	arr1, arr2 := make([]int, 0), make([]int, 0)
	for iter := len(num1) - 1; iter >= 0; iter-- {
		arr1 = append(arr1, int(num1[iter]-'0'))
	}
	for iter := len(num2) - 1; iter >= 0; iter-- {
		arr2 = append(arr2, int(num2[iter]-'0'))
	}

	//TODO 大数溢出问题

}
