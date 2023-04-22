package double_point

func isSubsequence(s string, t string) bool {
	sI := 0
	tI := 0
	for sI < len(s) && tI < len(t) {
		if s[sI] == t[tI] {
			sI++
			tI++
		} else {
			tI++
		}
	}
	return sI == len(s)
}
