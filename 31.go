package leetcode_offer

/**
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。


示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

提示：

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed是popped的排列。


*/

func ValidateStackSequences(pushed []int, popped []int) bool {
	var res []int
	//i := 0
	//pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
	//遍历pushed，直到遇到poped的第一位，说明该出栈了，
	for _, v := range pushed {
		res = append(res, v)
		for len(res) != 0 && res[len(res)-1] == popped[0] {
			popped = popped[1:]
			res = res[:len(res)-1]
		}
	}
	return len(res) == 0
}
