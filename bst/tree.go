// Package bst implements a binary search tree with arbitrary payloads.
package bst

func (t *Tree) Search(key int64) *Node {
	if t.Root == nil {
		return nil
	}

	return searchRecursive(t.Root, key)
}

func searchRecursive(node *Node, key int64) *Node {
	if node == nil || node.Key == key {
		return node
	}

	if key > node.Key {
		return searchRecursive(node.Right, key)
	}

	return searchRecursive(node.Left, key)
}
