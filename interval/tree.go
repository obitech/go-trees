package interval

import (
	"math"
	"sync"
)

type color int

const (
	red             color  = 0
	black           color  = 1
	sentinelPayload string = "sentinel"
)

type ErrNotFound string

func (e ErrNotFound) Error() string {
	return string(e)
}

// Tree represents an Interval tree with a root node and Mutex to
// protect concurrent access.
type Tree struct {
	lock     sync.RWMutex
	root     *node
	sentinel *node
}

// Result is a search result when looking up an interval in the tree.
type Result struct {
	Interval Interval
	Payload  interface{}
}

// NewIntervalTree returns an initialized but empty interval tree.
func NewIntervalTree() *Tree {
	sentinel := &node{color: black, payload: sentinelPayload}

	return &Tree{
		lock:     sync.RWMutex{},
		root:     sentinel,
		sentinel: sentinel,
	}
}

// Root returns a Result of the payload of the root node of the tree or an
// ErrNotFound if the tree is empty.
func (t *Tree) Root() (Result, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == t.sentinel {
		return Result{}, ErrNotFound("tree is empty")
	}

	return Result{
		Interval: t.root.key,
		Payload:  t.root.payload,
	}, nil
}

// Height returns the height (max depth) of the tree. Returns -1 if the tree
// has no nodes. A (rooted) tree with only a single node has a height of zero.
func (t *Tree) Height() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return int(t.height(t.root))
}

func (t *Tree) height(node *node) float64 {
	if node == t.sentinel {
		return -1
	}

	return 1 + math.Max(t.height(node.left), t.height(node.right))
}

// Min returns a Result of the lowest interval in the tree or an ErrNotFound if
// the tree is empty.
func (t *Tree) Min() (Result, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	n := t.min(t.root)

	if n == t.sentinel {
		return Result{}, ErrNotFound("tree is empty")
	}

	return Result{
		Interval: n.key,
		Payload:  n.payload,
	}, nil
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

	t.updateMax(x)
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

	t.updateMax(y)
}

func (t *Tree) newLeaf(key Interval, p interface{}) *node {
	return &node{
		key:     key,
		payload: p,
		left:    t.sentinel,
		right:   t.sentinel,
		max:     key.high,
	}
}

func (t *Tree) isLeaf(z *node) bool {
	return z.left == t.sentinel && z.right == t.sentinel
}

func (t *Tree) min(z *node) *node {
	for z != t.sentinel && z.left != t.sentinel {
		z = z.left
	}

	return z
}

func (t *Tree) updateMax(z *node) {
	z.max = z.key.high

	if z.right != t.sentinel && z.right.max.After(z.max) {
		z.max = z.right.max
	}

	if z.left != t.sentinel && z.left.max.After(z.max) {
		z.max = z.left.max
	}
}
