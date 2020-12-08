package leetcode_offer

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var dummy = &ListNode{}

	res := dummy
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			res.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		} else {
			res.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		}
		res = res.Next
	}
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}
	return dummy.Next

}
