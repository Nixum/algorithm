package string

// 进制
// 从高位到低位处理，起始让res=0，每次使用当前进位更新res
// 更新规则为 res = res * 10 + val
func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		c := int(columnTitle[i] - 'A' + 1)
		res = res * 26 + c
	}
	return res
}
