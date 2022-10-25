package graph

func equationsPossible(equations []string) bool {
	uf := InitUF(26)
	for _, str := range equations {
		if str[1] == '=' {
			uf.Union(int(str[0] - 'a'), int(str[3] - 'a'))
		}
	}
	for _, str := range equations {
		if str[1] == '!' {
			if uf.Connected(int(str[0] - 'a'), int(str[3] - 'a')) {
				return false
			}
		}
	}
	return true
}