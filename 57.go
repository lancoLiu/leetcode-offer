package leetcode_offer

var res1 [][]int

func findContinuousSequence(target int) [][]int {

	if target == 0 {
		return [][]int{}
	}
	var trackBack []int
	res1 = [][]int{}
	for i := 1; i <= target/2; i++ {
		trackBackContinuousSequence(target, 0, i, trackBack)
	}
	return res1
}

func trackBackContinuousSequence(target int, ans int, curr int, trackBack []int) {
	if ans > target {
		return
	}
	if ans == target {
		tmp := make([]int, len(trackBack))
		copy(tmp, trackBack)
		res1 = append(res1, append([]int{}, tmp...))
		return
	}
	trackBack = append(trackBack, curr)
	ans += curr
	curr++

	//next
	trackBackContinuousSequence(target, ans, curr, trackBack)

}
