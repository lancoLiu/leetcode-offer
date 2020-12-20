package leetcode_offer

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

*/

type MinStack struct {
	//最小栈时刻存储着栈顶元素存在时候的最小元素，在栈顶元素弹出的时候一同弹出最小栈栈顶元素
	minStack []int
	Stack    []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		minStack: nil,
		Stack:    nil,
	}
}

//push数据同时保存最小栈最小值
func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.minStack) > 0 && this.minStack[len(this.minStack)-1] <= x {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	} else {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		//TODO:
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return -1
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
