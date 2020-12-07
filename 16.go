package leetcode_offer

/**
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。


示例 1:

输入: 2.00000, 10
输出: 1024.00000

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	if n < 0 {
		x = 1 / x
		n = -n //golang int类型至少32位根据cpu判定
	}
	//直到n==0， n为1时候除2取整即为0余数1
	for n > 0 {
		if n&1 == 1 { //是否是奇数
			res *= x //奇数多出一位乘到结果中
		}
		x *= x
		n = n >> 1
	}
	return res
}
