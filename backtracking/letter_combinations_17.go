package backtracking

func letterCombinations(digits string) []string {
	resInLetterCombinations = make([]string, 0)
	pathInLetterCombinations = ""
	backTrackingInLetterCombinations(0, digits)
	return resInLetterCombinations
}

var resInLetterCombinations []string
var pathInLetterCombinations string
func backTrackingInLetterCombinations(start int, digits string) {
	if start > len(digits) - 1 {
		str := pathInLetterCombinations
		resInLetterCombinations = append(resInLetterCombinations, str)
		return
	}

	digit := digitMap[string(digits[start])]
	for i := 0; i < len(digit); i++ {
		pathInLetterCombinations += string(digit[i])
		backTrackingInLetterCombinations(start + 1, digits)
		pathInLetterCombinations = pathInLetterCombinations[:len(pathInLetterCombinations)-1]
	}
}

var digitMap = map[string]string{
	"0": "",
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}