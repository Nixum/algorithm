package string

import "math"

// 字符串合法的规则：
// 前缀只能是空格、+或者-，之后只能是数字，不允许连续出现+或-，否则输出0
// 不允许上下限溢出，溢出时取最值
func myAtoi(s string) int {
	start := 0
	// 去掉前缀空格
	for ; start < len(s); start++ {
		if s[start] != ' ' {
			break
		}
	}
	if start == len(s) {
		return 0
	}
	// 判断符号
	isNegative := false
	if s[start] == '+' {
		start++
	} else if s[start] == '-' {
		start++
		isNegative = true
	}
	res := 0
	for i := start; i < len(s); i++ {
		// 判断数字
		if s[i] > '9' || s[i] < '0' {
			break
		}
		// 判断溢出
		if res > math.MaxInt32 / 10 ||
			(res == math.MaxInt32 / 10 && (s[i] - '0') > math.MaxInt32 % 10) {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		// 注意要乘上符号
		res = res * 10 + int(s[i] - '0')
	}
	// 负数还是得在循环里面算，因为要判断是否溢出
	if isNegative {
		return -res
	}
	return res
}

// 这道题的edge case很多，解决不了溢出的问题
// 比如当s = "9223372036854775808"，要返回 2147483647
// 另外要注意，2的指数计算使用<<符号，^的意思是异或，同真异假
// 两个相同的数异或等于0，不同的数异或等于1
// 两个相同的数异或等于0，一个数和0异或等于它本身
func myAtoi1(s string) int {
	isNegative := false
	start := 0
	for ; start < len(s); start++ {
		if s[start] != ' ' {
			break
		}
	}
	end := len(s) - 1
	for ; end >= 0; end-- {
		if s[end] != ' ' {
			break
		}
	}
	resStr := ""
	for i := start; i <= end; i++ {
		if s[i] == '+' && i + 1 <= end && s[i + 1] == '-' {
			break
		}
		if s[i] == '+' {
			continue
		}
		if s[i] == '-' && i + 1 <= end && s[i + 1] - '0' <= 9 {
			isNegative = true
			continue
		}
		if s[i] - '0' > 9 {
			break
		}
		if s[i] - '0' <= 9 {
			numStart := i
			for i + 1 <= end {
				if s[i+1] - '0' <= 9 {
					i++
				} else {
					break
				}
			}
			resStr = s[numStart:i+1]
			break
		}
	}
	res := 0
	j := 0
	for j < len(resStr) {
		n := len(resStr) - j - 1
		num := 1
		for n > 0 {
			num *= 10
			n--
		}
		res += int(resStr[j] - '0') * num
		j++
	}
	if isNegative {
		res = -res
	}
	if res > 2<<30 - 1 {
		return 2<<30 - 1
	} else if res < -2<<30 {
		return -2<<30
	}
	return res
}