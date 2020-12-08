package leetcode_offer

import "container/list"

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

  4
 /  \
 2   7
/ \  / \
1  3 6  9
镜像输出：

  4
 /  \
 7   2
/ \  / \
9  6 3 1
限制：

0 <= 节点个数 <= 1000
*/
//递归处理,1,3交换，6，9交换；递归上层，2，7交换带动
func mirrorTree(root *TreeNode) *TreeNode {
	//终止条件： 当节点 rootroot 为空时（即越过叶节点），则返回 nullnull ；
	if root == nil {
		return root
	}

	//先暂存起右节点
	tmp := root.Right
	root.Right = mirrorTree(root.Left)

	root.Left = mirrorTree(tmp)
	return root

}

//迭代
/**
利用栈（或队列）遍历树的所有节点 nodenode ，并交换每个 nodenode 的左 / 右子节点
  4
 /  \
 2   7
/ \  / \
1  3 6  9
*/
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var queue = list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		e := queue.Front()
		node := e.Value.(*TreeNode)
		queue.Remove(e)
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
		if node.Right != nil {

			queue.PushBack(node.Right)

		}
		if node.Left != nil {
			queue.PushBack(node.Left)

		}
	}
	return root
}
