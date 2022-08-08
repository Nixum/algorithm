package common

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
