package double_point

// 空格替换成 %20
func replaceBlank(s string) string {
	ss := []byte(s)
	blankNum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			blankNum++
		}
	}
	res := make([]byte, len(ss) + 2 * blankNum)
	i := len(ss) - 1
	j := len(res) - 1
	for i >= 0 {
		if ss[i] != ' ' {
			res[j] = ss[i]
		} else {
			res[j] = '0'
			res[j - 1] = '2'
			res[j - 2] = '%'
			j = j - 2
		}
		j--
		i--
	}
	return string(res)
}
