package leetcode_offer

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
限制：

2 <= nums.length <= 10000

*/

//同样一个数组每个元素会出现两次，但是这个数组有两个元素只出现一次
//用异或能够消去偶数次相同的元素
func singleNumbers(nums []int) []int {

	res := 0
	for _, v := range nums {
		res ^= v
	}
	//res为只出现一次的元素a,b的亦或值
	//分组，在res的二进制表示中找出1，证明a和b这一位是不同的

	mask := 1
	for (res & mask) == 0 {
		mask <<= 1
	}
	a := 0
	b := 0
	for _, v1 := range nums {
		//运算符优先级
		if (mask & v1) == 0 {
			//划分到a
			a ^= v1
		} else {
			//划分到b
			b ^= v1
		}
	}
	return []int{a, b}

}
