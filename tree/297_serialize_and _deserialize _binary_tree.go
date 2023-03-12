package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	nodes []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	path := ""
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if node == nil {
			path += "#"
		} else {
			path += fmt.Sprintf("%d", node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
		if len(stack) > 0 {
			path += ","
		}
	}
	return path
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	c.nodes = strings.Split(data, ",")
	return c.CodecDeserialize()
}

func (c *Codec) CodecDeserialize() *TreeNode {
	if len(c.nodes) == 0 {
		return nil
	}
	node, nodes := CodecDeserializeNodesPop(c.nodes)
	c.nodes = nodes
	if node == "" || node == "#" {
		return nil
	}
	n, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   n,
		Left:  c.CodecDeserialize(),
		Right: c.CodecDeserialize(),
	}
	return root
}

func CodecDeserializeNodesPop(queue []string) (string, []string) {
	if len(queue) == 0 {
		return "", queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}

// 下面这套方法只适用于 根据下标能找到对应的节点那种场景
// 比如 [1,2,3,null,null,4,5,6,7] 并不是一个合法的二叉树
// 只适用 [1,2,3,null,null,4,5,null,null,null,null,6,7]
// Serializes a tree to a single string.
func (c *Codec) serialize2(root *TreeNode) string {
	if root == nil {
		return ""
	}
	path := ""
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			if node == nil {
				path += "#"
			} else {
				path += fmt.Sprintf("%d", node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			if i + 1 < size {
				path += ","
			}
		}
		if len(queue) > 0 {
			path += ","
		}
	}
	return path
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize2(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	queue := strings.Split(data, ",")
	n, _ := strconv.Atoi(string(data[0]))
	root := &TreeNode{
		Val: n,
	}
	CodecDeserialize2(root, 0, queue)
	return root
}

func CodecDeserialize2(node *TreeNode, i int, queue []string) {
	if node == nil {
		return
	}
	if 2 * i + 1 < len(queue) && queue[2 * i + 1] != "#" {
		n, _ := strconv.Atoi(queue[2 * i + 1])
		node.Left = &TreeNode{
			Val:   n,
		}
		CodecDeserialize2(node.Left, 2 * i + 1, queue)
	}
	if 2 * i + 2 < len(queue) && queue[2 * i + 2] != "#"{
		n, _ := strconv.Atoi(queue[2 * i + 2])
		node.Right = &TreeNode{
			Val:   n,
		}
		CodecDeserialize2(node.Right, 2 * i + 2, queue)
	}
}
