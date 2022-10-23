package bfs

func openLock(deadends []string, target string) int {
	lock := []byte{'0', '0', '0', '0'}
	res := 0
	deadEndMap := make(map[string]int)
	for i := range deadends {
		deadEndMap[deadends[i]] = 1
	}
	visited := make(map[string]int)
	visited[string(lock)] = 1
	queue := [][]byte{lock}
	var first []byte
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queuePop(queue)
			if len(first) == 0 {
				continue
			}
			if string(first) == target {
				return res
			}
			if _, exist := deadEndMap[string(first)]; exist {
				continue
			}
			// 每拨一次可能出现的结果
			for j := 0; j < 4; j++ {
				newLock1 := plus(first, j)
				if _, exist := visited[string(newLock1)]; !exist {
					visited[string(newLock1)] = 1
					queue = append(queue, newLock1)
				}
				newLock2 := dec(first, j)
				if _, exist := visited[string(newLock2)]; !exist {
					visited[string(newLock2)] = 1
					queue = append(queue, newLock2)
				}
			}
		}
		res++
	}
	return -1
}

func queuePop(queue [][]byte) ([]byte, [][]byte) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}

func plus(num []byte, i int) []byte {
	newNum := make([]byte, 0, 4)
	newNum = append(newNum, num...)
	if newNum[i] == '9' {
		newNum[i] = '0'
	} else {
		newNum[i]++
	}
	return newNum
}

func dec(num []byte, i int) []byte {
	newNum := make([]byte, 0, 4)
	newNum = append(newNum, num...)
	if newNum[i] == '0' {
		newNum[i] = '9'
	} else {
		newNum[i]--
	}
	return newNum
}