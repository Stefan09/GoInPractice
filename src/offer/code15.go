package offer

import (
	"math"
)

/**
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
例如，把 9 表示成二进制是 1001，有 2 位是 1。
因此，如果输入 9，则该函数输出 2。
*/

//左移易产生越界异常
func hammingWeight1(num uint32) int {
	if num > math.MaxUint32 {
		return -1
	}
	var counter = 0
	var iter uint32 = 1
	for i := 1; i <= 32; i++ {
		if num&iter != 0 {
			counter++
		}
		iter <<= 1
	}
	return counter
}

//右移
func HammingWeight2(num uint32) int {
	if num > math.MaxUint32 {
		return -1
	}
	counter := 0
	for num != 0 {
		if num%2 == 1 {
			counter++
		}
		num >>= 1
	}
	return counter
}
