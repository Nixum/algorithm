package double_point

// 思路：基于右括号的需求数，根据右括号的变化来判断是否要插入左括号
func minAddToMakeValid(s string) int {
	leftNeed := 0
	rightNeed := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rightNeed++
		}
		if s[i] == ')' {
			rightNeed--
			// 说明此时只有右括号
			if rightNeed == -1 {
				// 这个很关键
				rightNeed = 0
				leftNeed++
			}
		}
	}
	return leftNeed + rightNeed
}

// 错误的思路，只算了括号配对的场景，没有判断括号是否合法
func minAddToMakeValid2(s string) int {
	leftHad := 0
	rightHad := 0
	leftNeed := 0
	rightNeed := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftHad++
			if leftHad > rightHad {
				rightNeed++
			} else {
				rightNeed--
			}
		} else {
			rightHad++
			if rightHad > leftHad {
				leftNeed++
			} else {
				leftNeed--
			}
		}
	}
	return leftNeed + rightNeed
}
