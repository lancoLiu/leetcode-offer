package leetcode_offer

func movingCount(m int, n int, k int) int {
	res := make([][]bool, m)
	for i := 0; i < m; i++ {
		res[i] = make([]bool, n)
	}

	return movDfs(0, 0, m, n, k, res)
}

/**
递归终止条件：超出边界值;横纵坐标和超过k;是否计算过

*/
func movDfs(i, j, m, n, k int, res [][]bool) int {
	sum := calc(i) + calc(j)
	if i < 0 || i >= m || j < 0 || j >= n || sum > k || res[i][j] {
		return 0
	}
	res[i][j] = true
	return 1 + movDfs(i+1, j, m, n, k, res) + movDfs(i-1, j, m, n, k, res) + movDfs(i, j+1, m, n, k, res) + movDfs(i, j-1, m, n, k, res)
	//return 1 + movDfs(i+1,j,m,n,k,res)  + movDfs(i,j+1,m,n,k,res)
}

func calc(number int) int {
	res := 0
	for number != 0 {
		res += number % 10
		number = number / 10
	}
	return res
}
