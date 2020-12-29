package redblack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRBTree_Delete(t *testing.T) {
	t.Run("delete on empty tree is noop", func(t *testing.T) {
		tree := NewRedBlackTree()

		assert.Equal(t, -1, tree.Height())

		tree.Delete(15)

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("deleting unknown node on tree is noop", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(15, nil)
		tree.Upsert(20, nil)
		tree.Upsert(13, nil)

		assert.Equal(t, 1, tree.Height())

		tree.Delete(99)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, int64(15), tree.root.key)

		assert.Equal(t, red, tree.root.left.color)
		assert.Equal(t, int64(13), tree.root.left.key)

		assert.Equal(t, red, tree.root.right.color)
		assert.Equal(t, int64(20), tree.root.right.key)
	})

	t.Run("deleting root node leaves empty tree", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(15, nil)
		tree.Delete(15)

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("(h=1) delete right leaf is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(15, nil)
		tree.Upsert(20, nil)
		tree.Upsert(13, nil)

		tree.Delete(20)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, int64(15), tree.root.key)

		assert.Equal(t, red, tree.root.left.color)
		assert.Equal(t, int64(13), tree.root.left.key)

		assert.Equal(t, tree.sentinel, tree.root.right)
	})

	t.Run("(h=1) delete left leaf is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(15, nil)
		tree.Upsert(20, nil)
		tree.Upsert(13, nil)

		tree.Delete(13)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, int64(15), tree.root.key)

		assert.Equal(t, red, tree.root.right.color)
		assert.Equal(t, int64(20), tree.root.right.key)

		assert.Equal(t, tree.sentinel, tree.root.left)
	})

	t.Run("(h=3) delete from left subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []int64{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0} {
			tree.Upsert(i, i)
		}

		tree.Delete(1)
		tree.Delete(3)
		tree.Delete(5)
		tree.Delete(2)

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, int64(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.left
		assert.Equal(t, red, a.color)
		assert.Equal(t, int64(4), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, int64(0), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, int64(8), c.key)
	})

	t.Run("(h=3) delete from left subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []int64{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0} {
			tree.Upsert(i, i)
		}

		tree.Delete(5)
		tree.Delete(4)

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, int64(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.left
		assert.Equal(t, red, a.color)
		assert.Equal(t, int64(3), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, int64(1), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, int64(8), c.key)
	})

	t.Run("(h=3) delete from right subtree", func(t *testing.T) {
		tree := NewRedBlackTree()
		for _, i := range []int64{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30, 2, 0, 19, 15, 23} {
			tree.Upsert(i, i)
		}

		tree.Delete(20)
		tree.Delete(15)
		tree.Delete(13)
		tree.Delete(19)
		tree.Delete(30)
		tree.Delete(23)

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, int64(12), tree.root.key)
		assert.Equal(t, black, tree.root.color)

		a := tree.root.right
		assert.Equal(t, black, a.color)
		assert.Equal(t, int64(21), a.key)

		b := a.left
		assert.Equal(t, black, b.color)
		assert.Equal(t, int64(18), b.key)

		c := a.right
		assert.Equal(t, black, c.color)
		assert.Equal(t, int64(50), c.key)
	})
}
