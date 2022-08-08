package double_point

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	res := ""
	slow := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			res += string(s[i])
			slow++
		} else {
			if len(res) > 0 {
				res = res[0:len(res) - 1]
			}
		}
	}
	return res
}
