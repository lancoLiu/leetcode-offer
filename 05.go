package leetcode_offer

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。


示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

*/
func replaceSpace(s string) string {
	res := ""
	if len(s) == 0 {
		return res
	}
	for _, v := range s {
		if v == 32 {
			res += "%20"
		} else {
			res += string(v)
		}
	}
	return res

}
