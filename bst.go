package github.com/turelljj/turebst

import (
	"fmt"
)

// Bstree is an bst tree
type Bstree struct {
	root *node
}

//Insert inserts an value to an bst tree
func (b *Bstree) Insert(v int) *Bstree {
	if b.root == nil {
		b.root = &node{value: v, depth: 0}
		return b
	}
	b.root.insert(v)
	return b
}

// Sort return an bst tree elem in an slice
func (b *Bstree) Sort() (sorted []int) {
	if b.root == nil {
		return nil
	}
	sorted = append(sorted, b.root.sort()...)
	return sorted
}

// func (b *bstree) Remove(v int) *bstree {
// 	exsit, depth := b.root.search(v)
// 	if exsit {
// 	}
// 	return b
// }

// Find find the depth of given v in an bst tree
func (b *Bstree) Find(v int) (exsit bool, depth int) {
	if b.root == nil {
		return false, -1
	}
	if b.root.value == v {
		return true, 0
	}
	return b.root.search(v)
}

type node struct {
	left  *node // int lesser or equal to value
	right *node // int grater than value
	value int
	depth int
}

func (n *node) search(v int) (exsit bool, depth int) {
	if v < n.value && n.left != nil {
		return n.left.search(v)
	}
	if v > n.value && n.right != nil {
		return n.right.search(v)
	}
	if v == n.value {
		return true, n.depth
	}
	return false, -1
}

// PrintNode prints the node itslef and the subnode given an start node
func PrintNode(n *node) {
	fmt.Printf("%d: %v\n", n.depth, n.value)
	if n.left != nil {
		PrintNode(n.left)
	}
	if n.right != nil {
		PrintNode(n.right)
	}
}

func (n *node) sort() (sortedNode []int) {
	if n.left != nil {
		sortedNode = append(sortedNode, n.left.sort()...)
	}
	sortedNode = append(sortedNode, n.value)
	if n.right != nil {
		sortedNode = append(sortedNode, n.right.sort()...)
	}
	return sortedNode
}

func (n *node) insert(v int) *node {
	if n == nil {
		n = &node{}
	}
	if v <= n.value {
		if n.left == nil {
			n.left = &node{value: v, depth: n.depth + 1}
		} else {
			n.left.insert(v)
		}
	}
	if v > n.value {
		if n.right == nil {
			n.right = &node{value: v, depth: n.depth + 1}
		} else {
			n.right.insert(v)
		}
	}
	return n
}
