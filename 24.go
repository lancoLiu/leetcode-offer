package leetcode_offer

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//递归
func reverseList(head *ListNode) *ListNode {
	//这里必须注意是head.next；因为下面会head.Next.Next = head
	if head == nil || head.Next == nil {
		return head
	}

	//递归处理
	/**
	5->4->3->2<-1
	1
	*/
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead

}

//迭代
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//声明前节点
	var prev *ListNode
	//当前指向head节点下一个节点
	curr := head
	for head != nil {
		//保存head的next指向
		curr = head.Next
		//改变head的next指向
		head.Next = prev
		//令prev 指向head
		prev = head
		//head指向curr即原先自己的next，完成右移动
		head = curr
	}
	return prev
}
