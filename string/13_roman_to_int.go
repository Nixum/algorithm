package string

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		pre := 0
		if i > 0 {
			pre = romanMap[s[i-1]]
		}
		n := romanMap[s[i]]
		if pre < n {
			res += n-pre
		} else {
			res += n
		}
	}
	return res
}
