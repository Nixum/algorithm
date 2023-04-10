package string

import "strconv"

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber != 0 {
		// 因为这个26禁止不包含0
		// 比如正常x%26=[0,25], 但是从[1,26]取[A,Z]取模26%26=0不在范围内
		// 所以要-1，变成[0,25]
		columnNumber--
		c := columnNumber % 26
		res += strconv.Itoa('A' + c)
		columnNumber /= 26
	}
	str := ""
	for i := len(res)-1; i >= 0; i-- {
		str += string(res[i])
	}
	return str
}
