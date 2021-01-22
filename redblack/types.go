package redblack

import (
	"sync"
)

type color int

const (
	red   color = 0
	black color = 1
)

type Key interface {
	Less(k Key) bool
}

// Tree represents a red-black tree with a root node and Mutex to protect
// concurrent access.
type Tree struct {
	lock     sync.RWMutex
	root     *node
	sentinel *node
}

type node struct {
	key     Key
	color   color
	left    *node
	right   *node
	parent  *node
	payload interface{}
}
