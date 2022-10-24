package graph

type UF struct {
	count int
	parent []int
	size []int
}

func initUF(n int) *UF {
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

func (u *UF) Union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
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

func (u *UF) connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}