package greedy

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	nStr := fmt.Sprintf("%d", n)
	ss := []byte(nStr)
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i - 1] > ss[i] {
			ss[i - 1]--
			for j := i; j < len(ss); j++ {
				// 注意这里是字符
				ss[j] = '9'
			}
		}
	}
	resN, _ := strconv.Atoi(string(ss))
	return resN
}