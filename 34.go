package leetcode_offer

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。


[5,4,8,11,null,13,4,7,2,null,null,5,1]
22
示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//根据题意：从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
var res2 [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil && sum > 0 {
		return nil
	}
	res2 = [][]int{}
	recur3(root, sum, []int{})
	return res2
}

//target 还差多少凑齐；res结果集
func recur3(root *TreeNode, target int, res []int) {
	//递归终止条件
	if root == nil {
		return
	}
	res = append(res, root.Val)

	if root.Val == target && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,res 切片底层的数组会被修改
		//所以这里需要 copy res 到 tmp，再添加进 ret，防止 res 底层数据修改带来的错误
		tmp := make([]int, len(res))
		copy(tmp, res)
		res2 = append(res2, tmp)

	}

	recur3(root.Left, target-root.Val, res)
	recur3(root.Right, target-root.Val, res)
	//go中切片作为函数参数传递时，只是将切片的地址赋值给对应的形参，当函数内切片发生扩容时不会回传给原来的切片，也就是你的path参数其实每一次回到上一层就自动回退了
	res = res[:len(res)-1]
}
