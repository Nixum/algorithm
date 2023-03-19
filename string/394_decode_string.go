package string

import "strconv"

// 栈的解法, 用栈模拟递归
func decodeString(s string) string {
	res, multi, resStack, multiStack := "", 0, make([]string, 0), make([]int, 0)
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c >= '0' && c <= '9' {
			multi = multi * 10 + int(c-'0')
		} else if c == '[' {
			multiStack = append(multiStack, multi)
			resStack = append(resStack, res)
			multi = 0
			res = ""
		} else if c == ']'{
			tmp := ""
			curMulti := multiStack[len(multiStack)-1]
			multiStack = multiStack[:len(multiStack)-1]
			for i := 0; i < curMulti; i++ {
				tmp += res
			}
			curRes := resStack[len(resStack)-1]
			resStack = resStack[:len(resStack)-1]
			res = curRes + tmp
		} else {
			res += string(c)
		}
	}
	return res
}

// 递归解法
func decodeString1(s string) string {
	res, _ := dfsInDecodeString(s, 0)
	return res
}

func dfsInDecodeString(s string, start int) (string, int) {
	res, multi := "", 0
	for start < len(s) {
		c := s[start]
		if c >= '0' && c <= '9' {
			multi = multi * 10 + int(c-'0')
		} else if c == '[' {
			nextStr, nextIndex := dfsInDecodeString(s, start+1)
			for i := 0; i < multi; i++ {
				res += nextStr
			}
			start = nextIndex
		} else if c == ']'{
			return res, start+1
		} else {
			res += string(s[start])
		}
		start++
	}
	return res, start
}

// ============= 上面是下面的简洁版本 ==========
func decodeString2(s string) string {
	str := ""
	index := 0
	for index < len(s) {
		c := s[index]
		if c >= '0' && c <= '9' {
			numStr := ""
			for i := index; i < len(s); i++ {
				if s[i] >= '0' && s[i] <= '9' {
					numStr += string(s[i])
				} else {
					break
				}
			}
			index += len(numStr)
			n, _ := strconv.Atoi(numStr)
			start := index+1
			last := start+1
			leftNum := 1
			for i := start; i<len(s); i++ {
				if leftNum == 0 {
					break
				}
				if s[i] == '[' {
					leftNum++
				}
				if s[i] == ']' {
					leftNum--
				}
				last = i
			}
			tmp := ""
			nextStr := decodeString(s[start:last])
			for i := 0; i < int(n); i++ {
				tmp += nextStr
			}
			str += tmp
			index = last+1
		} else if c >= 'a' && c <= 'z' {
			str += string(s[index])
			index++
		}
	}
	return str
}
