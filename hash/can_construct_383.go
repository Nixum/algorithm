package hash

func canConstruct(r string, m string) bool {
	alphabet := make(map[int32]int)
	for _, c := range m {
		alphabet[c]++
	}
	for _, c := range r {
		count, exist := alphabet[c]
		if exist {
			count--
			if count < 0 {
				return false
			}
			alphabet[c] = count
		} else {
			return false
		}
	}
	return true
}

func canConstruct2(r string, m string) bool {
	tmp := make([]int, 26)
	for _, c := range m {
		tmp[c - 'a']++
	}
	for _, c := range r {
		tmp[c - 'a']--
		if tmp[c - 'a'] < 0 {
			return false
		}
	}
	return true
}
