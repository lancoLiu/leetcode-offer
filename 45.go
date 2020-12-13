package leetcode_offer

import (
	"fmt"
	"sort"
	"strings"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例2:

输入: [3,30,34,5,9]
输出: "3033459"
0 < nums.length <= 100

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
//先看每个数字的第一位，第一位相同的再去比较第二位，第二位的数小于等于第一位的（3，30）0小于3选择30，否则例如（3，34）
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})
	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
}

func compareNumber(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}
