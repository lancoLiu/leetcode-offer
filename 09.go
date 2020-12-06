package leetcode_offer

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )


示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对appendTail、deleteHead 进行10000次调用

*/

/**
先进先出
维护两个栈，第一个栈支持插入操作，第二个栈支持删除操作。
第二栈删除的操作：
如果为空，则将第一个栈的元素压入到第二个栈中
移除第二个栈的元素
*/
import "container/list"

type CQueue struct {
	s1 *list.List
	s2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

//队列尾部插入整数
//[1,2]
func (this *CQueue) AppendTail(value int) {
	//将元素插到末尾
	this.s1.PushBack(value)
}

//队列头部删除整数
func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.s2.Len() == 0 {
		for this.s1.Len() > 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	if this.s2.Len() != 0 {
		e := this.s2.Back()
		this.s2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
