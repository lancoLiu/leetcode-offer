package leetcode_offer

//哈希存储处理
func copyRandomList(head *Node) *Node {
	var mapsHelper = make(map[*Node]*Node)
	cur := head
	for cur != nil {
		temp := &Node{
			Val: cur.Val,
		}
		mapsHelper[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		mapsHelper[cur].Next = mapsHelper[cur.Next]
		mapsHelper[cur].Random = mapsHelper[cur.Random]
		cur = cur.Next
	}
	return mapsHelper[head]
}
