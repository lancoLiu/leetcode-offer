package leetcode_offer

import "strconv"

/**
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
动态规划
字母最大为25，思考最后一位一定是可以翻译的，但是最后两位能否翻译，必须认为nums[end-1,end]是否大于等于10小于等于25
所以dp[i]可以看作为nums[0..i]组成的个数
dp[i] =  dp[i - 1] + (dp[i - 2] && 10<=nums[i-2,i - 1] <=25)

*/
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		//加前一位
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if "10" <= pre && pre <= "25" {
			//加前两位
			r += p
		}

	}
	return r
}
