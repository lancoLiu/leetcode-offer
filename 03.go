package leetcode_offer

/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

*/
func findRepeatNumber(nums []int) int {
	var windows = make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := windows[v]; !ok {
			windows[v]++
		} else {
			return v
		}
	}
	return 0
}

//two
//原地置换，正常排序后，数字i应该在下标为i的位置
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
