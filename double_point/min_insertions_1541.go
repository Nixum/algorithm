package double_point

// 思路：基于右括号的需求数，根据右括号的变化来判断是否要插入左括号
func minInsertions(s string) int {
	// 需要插入的次数，可以是 ( , 也可以是 )
	insertNum := 0
	// 需要 ) 的次数
	rightNeed := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rightNeed += 2
			// 因为一个左括号要匹配两个右括号，此时右括号个数不能是奇数
			if rightNeed % 2 == 1 {
				insertNum++
				rightNeed--
			}
		}
		if s[i] == ')' {
			rightNeed--
			// 遇到一个多余的右括号
			if rightNeed == -1 {
				rightNeed = 1
				insertNum++
			}
		}
	}
	return insertNum + rightNeed
}
