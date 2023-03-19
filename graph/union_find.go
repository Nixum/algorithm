package graph

// 有向无环图，作用：最小生成树; 计算联通分量
// Kruskal 算法
// 判断两个节点是否在一个联通分量中
type UF struct {
	count int // 联通分量的数量
	parent []int
	size []int // 节点i下的节点数量
}

func InitUF(n int) *UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
		parent[i] = i
	}
	return &UF{
		count:  n,
		parent: parent,
		size:   size,
	}
}

// 将节点 p 和节点 q 连通
func (u *UF) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}
	if u.size[rootP] > u.size[rootQ] {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}
	u.count--
}

func (u *UF) Count() int {
	return u.count
}

// 判断节点 p 和节点 q 是否连通
func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// 返回节点 x 的连通分量根节点
func (u *UF) Find(x int) int {
	// 因为根节点肯定是指向自己的
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}