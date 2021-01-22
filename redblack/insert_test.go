package redblack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRBTree_Upsert(t *testing.T) {
	t.Run("insert on empty tree creates black root node", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), "test")

		assert.Equal(t, 0, tree.Height())
		assert.Equal(t, myInt(15), tree.root.key)
		assert.Equal(t, "test", tree.root.payload)
		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, "test", tree.Search(myInt(15)))
	})

	t.Run("upsert on existing root node key changes root node payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), "test")

		assert.Equal(t, 0, tree.Height())
		assert.Equal(t, myInt(15), tree.root.key)
		assert.Equal(t, "test", tree.root.payload)
		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, "test", tree.Search(myInt(15)))

		tree.Upsert(myInt(15), "test2")
		assert.Equal(t, 0, tree.Height())
		assert.Equal(t, myInt(15), tree.root.key)
		assert.Equal(t, "test2", tree.root.payload)
		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, "test2", tree.Search(myInt(15)))
	})

	t.Run("upsert to h=1 tree is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(15), "15")
		tree.Upsert(myInt(20), "20")
		tree.Upsert(myInt(13), "13")

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color, "root is black")
		assert.Equal(t, "15", tree.root.payload, "root is 15")
		assert.Equal(t, tree.sentinel, tree.root.parent)

		assert.Equal(t, red, tree.root.left.color, "left node is red")
		assert.Equal(t, "13", tree.root.left.payload, "left node is 13")
		assert.Equal(t, tree.root.left.parent, tree.root)

		assert.Equal(t, red, tree.root.right.color, "right node is red")
		assert.Equal(t, "20", tree.root.right.payload, "right node is 20")
		assert.Equal(t, tree.root.right.parent, tree.root)
	})

	t.Run("searching for 0 doesn't yield sentinel node", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(0), 0)
		tree.Upsert(myInt(15), nil)
		tree.Upsert(myInt(20), nil)
		tree.Upsert(myInt(3), nil)

		assert.Equal(t, 0, tree.Search(myInt(0)))
		assert.NotEqual(t, tree.sentinel, tree.Search(myInt(0)))
	})

	t.Run("(h=3) inserting into left-tree tree is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.Upsert(myInt(11), nil)
		tree.Upsert(myInt(2), nil)
		tree.Upsert(myInt(14), nil)
		tree.Upsert(myInt(15), nil)
		tree.Upsert(myInt(1), nil)
		tree.Upsert(myInt(7), nil)
		tree.Upsert(myInt(5), nil)
		tree.Upsert(myInt(8), nil)
		tree.Upsert(myInt(4), nil)

		root := tree.root
		rootL := tree.root.left
		rootR := tree.root.right
		rootLL := tree.root.left.left
		rootLR := tree.root.left.right
		rootLRL := tree.root.left.right.left
		rootRL := tree.root.right.left
		rootRR := tree.root.right.right
		rootRRR := tree.root.right.right.right

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, tree.sentinel, root.parent)
		assert.Equal(t, black, root.color)
		assert.Equal(t, myInt(7), root.key)

		assert.Equal(t, root, rootL.parent)
		assert.Equal(t, red, rootL.color)
		assert.Equal(t, myInt(2), rootL.key)

		assert.Equal(t, root, rootR.parent)
		assert.Equal(t, red, rootR.color)
		assert.Equal(t, myInt(11), rootR.key)

		assert.Equal(t, rootL, rootLL.parent)
		assert.Equal(t, black, rootLL.color)
		assert.Equal(t, myInt(1), rootLL.key)
		assert.True(t, tree.isLeaf(rootLL))

		assert.Equal(t, rootL, rootLR.parent)
		assert.Equal(t, black, rootLR.color)
		assert.Equal(t, myInt(5), rootLR.key)

		assert.Equal(t, rootLR, rootLRL.parent)
		assert.Equal(t, red, rootLRL.color)
		assert.Equal(t, myInt(4), rootLRL.key)
		assert.True(t, tree.isLeaf(rootLRL))

		assert.Equal(t, rootR, rootRL.parent)
		assert.Equal(t, black, rootRL.color)
		assert.Equal(t, myInt(8), rootRL.key)

		assert.Equal(t, rootR, rootRL.parent)
		assert.Equal(t, black, rootRL.color)
		assert.Equal(t, myInt(8), rootRL.key)
		assert.True(t, tree.isLeaf(rootRL))

		assert.Equal(t, rootR, rootRR.parent)
		assert.Equal(t, black, rootRR.color)
		assert.Equal(t, myInt(14), rootRR.key)

		assert.Equal(t, rootRR, rootRRR.parent)
		assert.Equal(t, red, rootRRR.color)
		assert.Equal(t, myInt(15), rootRRR.key)
		assert.True(t, tree.isLeaf(rootRRR))
	})

	t.Run("(h=3) inserting into right-tree is successful", func(t *testing.T) {
		tree := NewRedBlackTree()

		for _, i := range []int{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30} {
			tree.Upsert(myInt(i), i)
		}

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, black, tree.root.color)
		assert.Equal(t, myInt(12), tree.root.key)

		assert.Equal(t, red, tree.root.left.color)
		assert.Equal(t, myInt(3), tree.root.left.key)

		assert.Equal(t, black, tree.root.left.left.color)
		assert.Equal(t, myInt(1), tree.root.left.left.key)

		assert.Equal(t, black, tree.root.left.right.color)
		assert.Equal(t, myInt(5), tree.root.left.right.key)

		assert.Equal(t, red, tree.root.left.right.left.color)
		assert.Equal(t, myInt(4), tree.root.left.right.left.key)

		assert.Equal(t, red, tree.root.left.right.right.color)
		assert.Equal(t, myInt(8), tree.root.left.right.right.key)

		assert.Equal(t, red, tree.root.right.color)
		assert.Equal(t, myInt(20), tree.root.right.key)

		assert.Equal(t, black, tree.root.right.left.color)
		assert.Equal(t, myInt(18), tree.root.right.left.key)

		assert.Equal(t, red, tree.root.right.left.left.color)
		assert.Equal(t, myInt(13), tree.root.right.left.left.key)

		assert.Equal(t, black, tree.root.right.right.color)
		assert.Equal(t, myInt(30), tree.root.right.right.key)

		assert.Equal(t, red, tree.root.right.right.left.color)
		assert.Equal(t, myInt(21), tree.root.right.right.left.key)

		assert.Equal(t, red, tree.root.right.right.right.color)
		assert.Equal(t, myInt(50), tree.root.right.right.right.key)
	})
}

//
// func createTree(keys myInt) *Tree {
// 	rand.Seed(time.Now().UnixNano())
// 	tree := NewRedBlackTree()
//
// 	for i := int(1); i <= keys; i++ {
// 		tree.Upsert(rand.Int63n(keys), i)
// 	}
//
// 	return tree
// }
//
// var result interface{}
//
// func benchmarkSearch(i myInt, b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	tree := createTree(i)
//
// 	var r interface{}
//
// 	for n := 0; n < b.N; n++ {
// 		r = tree.Search(rand.Int63n(i))
// 	}
//
// 	result = r
// }
//
// func benchmarkDelete(i myInt, b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	tree := createTree(i)
//
// 	for n := 0; n < b.N; n++ {
// 		tree.Delete(rand.Int63n(i))
// 	}
// }
//
// func BenchmarkRBTree_Upsert(b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	tree := NewRedBlackTree()
//
// 	for n := 1; n <= b.N; n++ {
// 		tree.Upsert(rand.Int63n(myInt(n)), nil)
// 	}
// }
//
// func BenchmarkRBTree_Search10_000(b *testing.B) {
// 	benchmarkSearch(10_000, b)
// }
//
// func BenchmarkRBTree_Search100_000(b *testing.B) {
// 	benchmarkSearch(100_000, b)
// }
//
// func BenchmarkRBTree_Search1_000_000(b *testing.B) {
// 	benchmarkSearch(1_000_000, b)
// }
//
// func BenchmarkRBTree_Delete10_000(b *testing.B) {
// 	benchmarkDelete(10_000, b)
// }
//
// func BenchmarkRBTree_Delete100_000(b *testing.B) {
// 	benchmarkDelete(100_000, b)
// }
//
// func BenchmarkRBTree_Delete1_000_000(b *testing.B) {
// 	benchmarkDelete(1_000_000, b)
// }
