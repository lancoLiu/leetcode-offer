package leetcode_offer

/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

*/

type MaxQueue struct {
	queue       []int
	maxValQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:       []int{},
		maxValQueue: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxValQueue) == 0 {
		return -1
	}
	//返回最开始的那个
	return this.maxValQueue[0]
}

//往队列入元素，较大值进队，较小值出队，用于存储最大值，呈递减序列
func (this *MaxQueue) Push_back(value int) {
	//插入队列1
	this.queue = append(this.queue, value)
	//注意是for
	for len(this.maxValQueue) != 0 && value >= this.maxValQueue[len(this.maxValQueue)-1] {
		this.maxValQueue = this.maxValQueue[:len(this.maxValQueue)-1]
	}
	this.maxValQueue = append(this.maxValQueue, value)
}

//从队列开头拿元素;队列的首部与 queue 出队的值作比较，如果相同，那么一起出队
func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	node := this.queue[0]
	this.queue = this.queue[1:]
	if this.maxValQueue[0] == node {
		this.maxValQueue = this.maxValQueue[1:]
	}
	return node
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
