package leetcode_offer

/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0

[1,2,3,4,5]
[3,4,5,1,2]
[2,2,2,0,1]
根据上面可得必须的mid与right去做比较
为了能比较，right必须是number里面的，所以初始化为len-1
当mid>r 说明mid肯定不是最小值，mid后面是旋转点，所以让l=mid+1
当mid == r 无法确定旋转点在哪，当 x < j 时： 易得执行 j = j - 1 后，旋转点 x 仍在区间 [i, j] 内。
当 x = j 时： 执行 j = j - 1 后越过（丢失）了旋转点 x ，但最终返回的元素值 nums[i] 仍等于旋转点值 nums[x]
当mid < r 旋转点一定在[l,mid]中，所以缩减范围为[l,mid]
*/

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			//[mid+1,r]
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			//[l,mid]
			r = mid
		} else {
			//[l,r-1]
			r--
		}
	}
	return numbers[l]
}
