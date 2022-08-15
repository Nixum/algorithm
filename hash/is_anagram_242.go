package hash

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	for _, c := range s {
		alphabet[c - 'a']++
	}
	for _, c := range t {
		alphabet[c - 'a']--
	}
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}
