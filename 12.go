package leetcode_offer

/**
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，
每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true


递归查找

*/

func exist(board [][]byte, word string) bool {

	rows := len(board)
	col := len(board[0])

	for i := 0; i < rows; i++ {

		for j := 0; j < col; j++ {
			if dfs2(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false

}

/**
递归停止条件;越界;当前字符串不等于要查找的;
*/
func dfs2(board [][]byte, word string, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if len(word)-1 == k {
		return true
	}
	tmp := board[i][j]
	board[i][j] = 0
	res := dfs2(board, word, i-1, j, k+1) || dfs2(board, word, i+1, j, k+1) || dfs2(board, word, i, j-1, k+1) || dfs2(board, word, i, j+1, k+1)
	board[i][j] = tmp
	return res
}
