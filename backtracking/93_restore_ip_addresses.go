package backtracking

import (
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	resInRestoreIpAddresses = make([]string, 0)
	pathInRestoreIpAddresses = ""
	backTrackingInRestoreIpAddresses(s, 0, 0)
	return resInRestoreIpAddresses
}

var resInRestoreIpAddresses []string
var pathInRestoreIpAddresses string
func backTrackingInRestoreIpAddresses(s string, start int, point int) {
	if point == 3 {
		t := s[start:]
		if isValidIpInRestoreIpAddresses(t) {
			tmp := pathInRestoreIpAddresses
			tmp += t
			resInRestoreIpAddresses = append(resInRestoreIpAddresses, tmp)
		}
		return
	}
	for i := start; i < len(s); i++ {
		t := s[start : i + 1]
		if !isValidIpInRestoreIpAddresses(t) {
			break
		}
		pathInRestoreIpAddresses += t + "."
		backTrackingInRestoreIpAddresses(s, i + 1, point + 1)
		pathInRestoreIpAddresses = pathInRestoreIpAddresses[:len(pathInRestoreIpAddresses) - len(t) - 1]
	}
}

func isValidIpInRestoreIpAddresses(str string) bool {
	if str == "" {
		return false
	}
	if strings.HasPrefix(str, "0") && len(str) > 1{
		return false
	}
	if len(str) > 3 {
		return false
	}
	multi := 1
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		sum += int(str[i] - '0') * multi
		multi *= 10
	}
	if sum > 255 {
		return false
	}
	return true
}

var pathInRestoreIpAddresses2 []string
func backTrackingInRestoreIpAddresses2(s string, start int) {
	if len(pathInRestoreIpAddresses2) == 3 {
		path := s[start: ]
		if isValidIpInRestoreIpAddresses(path) {
			// key point，此时path不能加入pathInRestoreIpAddresses2，否则会导致重复
			// 因为这里加完，返回上层再移除，再上层会重复加入path
			resInRestoreIpAddresses = append(resInRestoreIpAddresses, fmt.Sprintf("%s.%s.%s.%s", pathInRestoreIpAddresses2[0], pathInRestoreIpAddresses2[1], pathInRestoreIpAddresses2[2], path))
		}
		return
	}

	for i := start; i < len(s); i++ {
		path := string(s[start : i + 1])
		if !isValidIpInRestoreIpAddresses(path) {
			break
		}
		pathInRestoreIpAddresses2 = append(pathInRestoreIpAddresses2, path)
		backTrackingInRestoreIpAddresses2(s, i + 1)
		pathInRestoreIpAddresses2 = pathInRestoreIpAddresses2[: len(pathInRestoreIpAddresses2) - 1]
	}
}