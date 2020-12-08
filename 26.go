package leetcode_offer

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

  3
  / \
 4  5
 / \
1  2
给定的树 B：

    4
 /
1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

限制：

0 <= 节点个数 <= 10000
*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	//初始化
	if A == nil || B == nil {
		return false
	}
	//在A的左子树或者右子树找到和B一样的子树即可
	return findSameTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

//递归查找
func findSameTree(root, sameTree *TreeNode) bool {

	//终止条件
	if sameTree == nil {
		return true
	}
	if root == nil || root.Val != sameTree.Val {
		return false
	}

	//相等
	return findSameTree(root.Left, sameTree.Left) && findSameTree(root.Right, sameTree.Right)

}
