package leetcode_offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/
/**
先打印存储，再反转切片即可
*/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	j := len(res) - 1
	//1,3,2,4
	//4,2,3,1
	for i := 0; i < len(res)/2; i++ {
		res[i], res[j] = res[j], res[i]
		j--
	}
	return res
}
