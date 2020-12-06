package leetcode_offer

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000


*/

//输入前序和中序，获得树,使用递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIndex := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
	return root
}

//迭代
func buildTree2(preorder []int, inorder []int) *TreeNode {
	//找到根结点
	//左边为根结点左子树，右边为右子树
	if len(preorder) == 0 {
		return nil
	}
	//建立树根
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	length := len(preorder)
	stack2 := Stack{root}

	inorderIndex := 0
	//从1开始遍历前序遍历的每个元素
	for i := 1; i < length; i++ {
		preorderVal := preorder[i]
		node := stack2.Peek()
		//判断其上一个元素（即栈顶元素）是否等于中序遍历下标指向的元素。

		//若上一个元素不等于中序遍历下标指向的元素，则将当前元素作为其上一个元素的左子节点，并将当前元素压入栈内。
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{
				preorderVal,
				nil,
				nil,
			}
			//入栈
			stack2.Push(node.Left)
		} else {
			//若上一个元素等于中序遍历下标指向的元素，则从栈内弹出一个元素，同时令中序遍历下标指向下一个元素，之后继续判断栈顶元素是否等于中序遍历下标指向的元素，若相等则重复该操作，直至栈为空或者元素不相等。然后令当前元素为最后一个想等元素的右节点。
			for len(stack2) != 0 && stack2.Peek().Val == inorder[inorderIndex] {
				node = stack2.Pop()
				inorderIndex++
			}
			node.Right = &TreeNode{
				preorderVal,
				nil,
				nil,
			}
			stack2.Push(node.Right)
		}
	}
	return root
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *Stack) Peek() *TreeNode {
	if len(*s) == 0 {
		return nil
	}
	n := (*s)[len(*s)-1]
	return n
}
