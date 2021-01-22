// Package redblack implements a Red-Black tree, which is a balanced binary
// search tree that runs on O(lg n) on all operations.
package redblack

import (
	"math"
	"sync"
)

// NewRedBlackTree returns a new red-back tree. All operations on the tree are
// safe to be accessed concurrently.
func NewRedBlackTree() *Tree {
	sentinel := &node{color: black, payload: "sentinel"}

	return &Tree{
		lock:     sync.RWMutex{},
		root:     sentinel,
		sentinel: sentinel,
	}
}

// Root returns the payload of the root node of the tree.
func (t *Tree) Root() interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == t.sentinel {
		return nil
	}

	return t.root.payload
}

// Height returns the height (max depth) of the tree. Returns -1 if the tree
// has no nodes. A (rooted) tree with only a single node has a height of zero.
func (t *Tree) Height() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return int(t.height(t.root))
}

// Min returns the payload of the lowest key, or nil.
func (t *Tree) Min() interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	n := t.min(t.root)

	if n == t.sentinel {
		return nil
	}

	return n.payload
}

// Max returns the payload of the highest key, or nil.
func (t *Tree) Max() interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	n := t.max(t.root)

	if n == t.sentinel {
		return nil
	}

	return n.payload
}

// Search returns the payload for a given key, or nil.
func (t *Tree) Search(key Key) interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == nil {
		return nil
	}

	n := t.search(t.root, key)

	if n == t.sentinel {
		return nil
	}

	return n.payload
}

// Successor returns the payload of the next highest neighbour (key-wise) of the
// passed key.
func (t *Tree) Successor(key Key) interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	n := t.successor(t.search(t.root, key))

	if n == nil {
		return nil
	}

	return n.payload
}

func (t *Tree) height(node *node) float64 {
	if node == t.sentinel {
		return -1
	}

	return 1 + math.Max(t.height(node.left), t.height(node.right))
}

func (t *Tree) successor(z *node) *node {
	if z == t.sentinel {
		return nil
	}

	if z.right != t.sentinel {
		return t.min(z.right)
	}

	parent := z.parent

	for parent != t.sentinel && z == parent.right {
		z = parent
		parent = z.parent
	}

	return parent
}

func (t *Tree) min(z *node) *node {
	for z != t.sentinel && z.left != t.sentinel {
		z = z.left
	}

	return z
}

func (t *Tree) max(z *node) *node {
	for z != t.sentinel && z.right != t.sentinel {
		z = z.right
	}

	return z
}

func (t *Tree) search(z *node, key Key) *node {
	if z == t.sentinel || z.key == key {
		return z
	}

	for z != t.sentinel && z.key != key {
		if z.key.Less(key) {
			z = z.right
		} else {
			z = z.left
		}
	}

	return z
}

func (t *Tree) rotateLeft(x *node) {
	// y's left subtree will be x's right subtree.
	y := x.right
	x.right = y.left

	if y.left != t.sentinel {
		y.left.parent = x
	}

	// Restore parent relationships.
	y.parent = x.parent

	switch {
	case x.parent == t.sentinel:
		t.root = y
	case x.parent.left == x:
		x.parent.left = y
	default:
		x.parent.right = y
	}

	// x will be y's new left-child.
	y.left = x
	x.parent = y
}

func (t *Tree) rotateRight(x *node) {
	y := x.left
	x.left = y.right

	if y.right != t.sentinel {
		y.right.parent = x
	}

	y.parent = x.parent

	switch {
	case x.parent == t.sentinel:
		t.root = y
	case x.parent.left == x:
		x.parent.left = y
	default:
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (t *Tree) newLeaf(key Key, p interface{}) *node {
	return &node{
		key:     key,
		payload: p,
		left:    t.sentinel,
		right:   t.sentinel,
	}
}

func (t *Tree) isLeaf(z *node) bool {
	return z.left == t.sentinel && z.right == t.sentinel
}
