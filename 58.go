package leetcode_offer

import "strings"

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。



示例 1：

输入: "the sky is blue"
输出:"blue is sky the"
示例 2：

输入: " hello world! "
输出:"world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
双指针
初始化指向尾部
移动l直到找到第一个单词左边界，添加
移动l找到倒数第二个单词右边界
同步r

*/
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.TrimSpace(s)
	l, r := len(s)-1, len(s)-1
	str := ""
	for l >= 0 {
		for l >= 0 && s[l] != ' ' {
			l--
		}
		str += s[l+1:r+1] + " "
		for l >= 0 && s[l] == ' ' {
			l--
		}
		r = l

	}
	return strings.TrimSpace(str)
}
