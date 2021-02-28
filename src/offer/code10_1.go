package offer

/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

//递归
func Fib1(n int) int {
	if n < 0 || n > 100 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	res := Fib1(n-1) + Fib1(n-2)
	return res % int(1e9+7)
}

//迭代
func Fib2(n int) int {
	if n < 0 || n > 100 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	pre1, pre2 := 0, 1
	for i := 2; i <= n; i++ {
		pre1, pre2 = pre2%int(1e9+7), (pre1+pre2)%int(1e9+7)
	}
	return pre2
}
