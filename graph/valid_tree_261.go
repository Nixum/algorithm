package graph

// 判断输入的若干条边是否能构造出一棵树
func validTree(n int, edges [][]int) bool {
	uf := InitUF(n)
	for _, edge := range edges {
		if uf.Connected(edge[0], edge[1]) {
			return false
		}
		uf.Union(edge[0], edge[1])
	}
	return uf.Count() == 1
}
