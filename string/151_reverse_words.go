package string

// 最优雅写法
func reverseWords(s string) string {
	ss := []byte(s)
	slow := 0
	fast := 0
	// 移除前置空格
	for fast < len(ss) && ss[fast] == ' ' {
		fast++
	}
	// 移除中间空格
	for ; fast < len(ss); fast++ {
		if fast - 1 > 0 && ss[fast] == ' ' && ss[fast] == ss[fast - 1] {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}
	// 移除后置空格
	if slow - 1 > 0 && ss[slow - 1] == ' ' {
		ss = ss[:slow - 1]
	} else {
		ss = ss[:slow]
	}
	reverseWordsReverse(ss, 0, len(ss) - 1)
	i := 0
	for i < len(ss) {
		j := i
		for j < len(ss) && ss[j] != ' '{
			j++
		}
		reverseWordsReverse(ss, i, j - 1)
		i = j + 1
	}
	return string(ss)
}

// 还不够优雅
func reverseWords2(s string) string {
	ss := []byte(s)
	// 移除前置空格
	fast := 0
	for fast < len(ss) && ss[fast] == ' ' {
		fast++
	}
	end := len(ss) - 1
	// 移除后面空格
	for end >= 0 && ss[end] == ' ' {
		end--
	}
	// 移除中间空格
	index := 0
	for i := fast; i <= end; i++ {
		if i - 1 >= 0 && ss[i] == ss[i - 1] && ss[i] == ' ' {
			continue
		}
		ss[index] = ss[i]
		index++
	}
	ss = ss[:index]
	left := 0
	right := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			right = i - 1
			reverseWordsReverse(ss, left, right)
			left = i + 1
		}
		if i + 1 == len(ss) {
			right = i
			reverseWordsReverse(ss, left, right)
			left = i + 1
		}
	}
	res := make([]byte, 0)
	for i := len(ss) - 1; i >= 0; i-- {
		res = append(res, ss[i])
	}
	return string(res)
}

// 没有解决去除多余空格的问题
func reverseWords3(s string) string {
	ss := []byte(s)
	left := 0
	right := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			right = i - 1
			reverseWordsReverse(ss, left, right)
			left = i + 1
		}
		if i + 1 == len(ss) {
			right = i
			reverseWordsReverse(ss, left, right)
			left = i + 1
		}
	}
	res := make([]byte, 0)
	for i := len(ss) - 1; i >= 0; i-- {
		res = append(res, ss[i])
	}
	return string(res)
}

func reverseWordsReverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
