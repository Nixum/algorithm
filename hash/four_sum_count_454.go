package hash

func fourSumCount(a []int, b []int, c []int, d []int) int {
	m := make(map[int]int)
	for _, i := range a {
		for _, j := range b {
			m[i+j]++
		}
	}
	count := 0
	for _, i := range c {
		for _, j := range d {
			if c, exist := m[0 - (i + j)]; exist {
				count += c
			}
		}
	}
	return count
}
