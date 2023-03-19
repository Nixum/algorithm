package tree

// 前缀树
type Trie struct {
	root *PrefixTreeNode
}

type PrefixTreeNode struct {
	// 这个是关键，避免相同前缀的单词insert后，查询出错
	// 比如 apple先插入了，查询apple、前缀查询app、查询app 时做区分
	isEnd bool
	Set map[byte]*PrefixTreeNode
}

func PrefixTreeConstructor() Trie {
	return Trie{
		root: &PrefixTreeNode{
			isEnd: false,
			Set: make(map[byte]*PrefixTreeNode, 0),
		},
	}
}


func (t *Trie) Insert(word string)  {
	root := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if child, exist := root.Set[c]; !exist {
			root.Set[c] = &PrefixTreeNode{
				Set:      make(map[byte]*PrefixTreeNode, 0),
			}
			root = root.Set[c]
		} else {
			root = child
		}
	}
	root.isEnd = true
}


func (t *Trie) Search(word string) bool {
	root := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if child, exist := root.Set[c]; exist {
			root = child
		} else {
			return false
		}
	}
	return root.isEnd
}


func (t *Trie) StartsWith(prefix string) bool {
	root := t.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if child, exist := root.Set[c]; exist {
			root = child
		} else {
			return false
		}
	}
	return true
}