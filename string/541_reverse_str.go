package string

// 每隔2k的前k个进行翻转
func reverseStr(s string, k int) string {
	res := ""
	for i := 0; i < len(s); i = i + 2 * k {
		end := i + k
		if end > len(s) {
			end = len(s)
			res += string(reverseStrDetail([]byte(s[i:end])))
		} else {
			end2 := i + 2*k
			if end2 > len(s) {
				end2 = len(s)
			}
			res += string(reverseStrDetail([]byte(s[i:end]))) + s[end:end2]
		}
	}
	return res
}

func reverseStr2(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(ss); i = i + 2 * k {
		if i + k <= len(ss) {
			reverseStrDetail(ss[i : i + k])
			continue
		}
		reverseStrDetail(ss[i:len(ss)])
	}
	return string(ss)
}

func reverseStrDetail(s []byte) []byte {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

