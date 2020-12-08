package leetcode_offer

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

  1
 / \
 2  2
/ \ / \
3 4 4 3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

  1
 / \
 2  2
 \  \
 3  3
0 <= 节点个数 <= 1000
*/
//递归，判断二叉树的左子树和右子树是否对称，即左子树的左节点与右子树右节点；左子树右节点与左子树左节点
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(a, b *TreeNode) bool {
	//递归终止条件
	//如果都到叶子结点遍历完，说明相同
	if a == nil && b == nil {
		return true
	}
	//否则某一个已经到叶子结点不同，false
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	//递归判断
	return recur(a.Left, b.Right) && recur(a.Right, b.Left)
}
