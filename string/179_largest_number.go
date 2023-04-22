package string

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	numStrs := make([]string, 0, len(nums))
	for _, n := range nums {
		numStrs = append(numStrs, strconv.Itoa(n))
	}
	// 因为 s1 和 s2 是等长的，字符串比较是按位比较，所以转成数字也是成立
	sort.Slice(numStrs, func(i, j int) bool {
		s1 := numStrs[i] + numStrs[j]
		s2 := numStrs[j] + numStrs[i]
		return s1 > s2
	})
	res := ""
	for _, str := range numStrs {
		res += str
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
