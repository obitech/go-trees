// Package bst implements a binary search tree with arbitrary payloads.
package bst

// Search returns a Node based on a key, or nil.
func (t *Tree) Search(key int64) *Node {
	if t.Root == nil {
		return nil
	}

	return searchRecursive(t.Root, key)
}

// Min returns the Node with the lowest key, or nil.
func (t *Tree) Min() *Node {
	return min(t.Root)
}

func min(node *Node) *Node {
	for node != nil && node.Left != nil {
		node = node.Left
	}

	return node
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

func searchIterative(node *Node, key int64) *Node {
	if node == nil || node.Key == key {
		return node
	}

	for node != nil && node.Key != key {
		if key > node.Key {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	return node
}
