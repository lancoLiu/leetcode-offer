package leetcode_offer

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [
[1,2,3],
[4,5,6],
[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	//左右上下边界
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var res []int
	for {
		//打印第一行
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if b < t {
			break
		}

		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return res

}
