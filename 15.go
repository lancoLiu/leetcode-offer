package leetcode_offer

/**
请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*/

//二进制运算
/**
数字本身a和这个减1的b的与运算能消去a的最后一个二进制1
*/
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

/**
若 n & 1 = 0 ，则 n 二进制 最右一位 为 0
若 n & 1 = 1 ，则 n 二进制 最右一位 为 1
*/
func hammingWeight2(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
