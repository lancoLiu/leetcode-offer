package leetcode_offer

/**

斐波那契数列
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % 1000000007
	}
	return res[n]
}
