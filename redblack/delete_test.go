package redblack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type myInt int

func (i myInt) Less(v Key) bool {
	return i < v.(myInt)
}

func TestRBTree_Delete(t *testing.T) {
	t.Run("delete on empty tree is noop", func(t *testing.T) {
		tree := NewRedBlackTree()

		assert.Equal(t, -1, tree.Height())

		tree.Delete(myInt(15))

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("deleting unknown node on tree is noop", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), nil)
		tree.Upsert(myInt(20), nil)
		tree.Upsert(myInt(13), nil)

		assert.Equal(t, 1, tree.Height())

		tree.Delete(myInt(99))

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, myInt(15), tree.root.key)

		assert.Equal(t, red, tree.root.left.color)
		assert.Equal(t, myInt(13), tree.root.left.key)

		assert.Equal(t, red, tree.root.right.color)
		assert.Equal(t, myInt(20), tree.root.right.key)
	})

	t.Run("deleting root node leaves empty tree", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), nil)
		tree.Delete(myInt(15))

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("(h=1) delete right leaf is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), nil)
		tree.Upsert(myInt(20), nil)
		tree.Upsert(myInt(13), nil)

		tree.Delete(myInt(20))

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, myInt(15), tree.root.key)

		assert.Equal(t, red, tree.root.left.color)
		assert.Equal(t, myInt(13), tree.root.left.key)

		assert.Equal(t, tree.sentinel, tree.root.right)
	})

	t.Run("(h=1) delete left leaf is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), nil)
		tree.Upsert(myInt(20), nil)
		tree.Upsert(myInt(13), nil)

		tree.Delete(myInt(13))

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, myInt(15), tree.root.key)

		assert.Equal(t, red, tree.root.right.color)
		assert.Equal(t, myInt(20), tree.root.right.key)

		assert.Equal(t, tree.sentinel, tree.root.left)
	})

	t.Run("(h=3) delete from left subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []int{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0} {
			tree.Upsert(myInt(i), i)
		}

		tree.Delete(myInt(1))
		tree.Delete(myInt(3))
		tree.Delete(myInt(5))
		tree.Delete(myInt(2))

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, myInt(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.left
		assert.Equal(t, red, a.color)
		assert.Equal(t, myInt(4), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, myInt(0), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, myInt(8), c.key)
	})

	t.Run("(h=3) delete from left subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []myInt{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0} {
			tree.Upsert(myInt(i), i)
		}

		tree.Delete(myInt(5))
		tree.Delete(myInt(4))

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, myInt(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.left
		assert.Equal(t, red, a.color)
		assert.Equal(t, myInt(3), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, myInt(1), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, myInt(8), c.key)
	})

	t.Run("(h=3) delete from right subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []myInt{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0, 19, 15, 23} {
			tree.Upsert(myInt(i), i)
		}

		tree.Delete(myInt(20))
		tree.Delete(myInt(15))
		tree.Delete(myInt(13))
		tree.Delete(myInt(19))
		tree.Delete(myInt(30))
		tree.Delete(myInt(23))

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, myInt(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.right
		assert.Equal(t, black, a.color)
		assert.Equal(t, myInt(21), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, myInt(18), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, myInt(50), c.key)
	})
}
