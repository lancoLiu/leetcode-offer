package leetcode_offer

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
数组长度 <= 1000
*/

//递归处理
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	//递归，左子树小于跟节点，右子树大于根节点
	//分置递归，找到之后判断左子树和右子树的左右子树区间是否满足
	return recur2(postorder, 0, len(postorder)-1)
}

func recur2(postorder []int, l, r int) bool {
	//递归终止条件
	if l >= r {
		return true
	}
	mid := l
	//找到mid使得区间[l,mid - 1]是左子树
	for postorder[mid] < postorder[r] {
		mid++
	}

	//还要判断[mid,r]是否都满足大于postdata[r]
	for i := mid; i < r; i++ {
		if postorder[i] < postorder[r] {
			return false
		}

	}
	return recur2(postorder, l, mid-1) && recur2(postorder, mid, r-1)
}

//TODO:单调栈
