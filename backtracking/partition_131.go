package backtracking

func partition(str string) [][]string {
	resInPartition = make([][]string, 0)
	pathInPartition = make([]string, 0)
	backTrackingInPartition(str, 0)
	return resInPartition
}

var resInPartition [][]string
var pathInPartition []string
func backTrackingInPartition(str string, start int) {
	if start >= len(str) {
		res := make([]string, 0)
		res = append(res, pathInPartition...)
		resInPartition = append(resInPartition, res)
		return
	}
	path := ""
	for i := start; i < len(str); i++ {
		path += string(str[i])
		if !isPartitionInPartition(path) {
			continue
		}
		pathInPartition = append(pathInPartition, path)
		backTrackingInPartition(str, i + 1)
		pathInPartition = pathInPartition[:len(pathInPartition) - 1]
	}
}

func isPartitionInPartition(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
