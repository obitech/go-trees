package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type item struct {
	key     int64
	payload string
}

func TestTree_search(t *testing.T) {
	tt := []struct {
		name string
		tree *BSTree
		key  int64
		want *node
	}{
		{
			name: "nil tree returns nil",
			tree: &BSTree{},
		},
		{
			name: "root hit returns value",
			tree: &BSTree{
				root: &node{
					key: 5,
				},
			},
			key:  5,
			want: &node{key: 5},
		},
		{
			name: "root miss returns nil",
			tree: &BSTree{
				root: &node{
					key: 6,
				},
			},
			key: 5,
		},
		{
			name: "h=1 left tree hit returns value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 15,
					},
				},
			},
			key:  5,
			want: &node{key: 5},
		},
		{
			name: "h=2 right tree hit returns value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 15,
					},
				},
			},
			key:  15,
			want: &node{key: 15},
		},
		{
			name: "h=2 miss returns nil",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 15,
					},
				},
			},
			key: 99,
		},
		{
			name: "h=3 right tree hit returns value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
						left: &node{
							key: 3,
						},
						right: &node{
							key: 7,
						},
					},
					right: &node{
						key: 15,
						left: &node{
							key: 12,
						},
						right: &node{
							key: 19,
						},
					},
				},
			},
			key:  19,
			want: &node{key: 19},
		},
		{
			name: "h=3 left tree hit returns value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
						left: &node{
							key: 3,
						},
						right: &node{
							key: 7,
						},
					},
					right: &node{
						key: 15,
						left: &node{
							key: 12,
						},
						right: &node{
							key: 19,
						},
					},
				},
			},
			key:  3,
			want: &node{key: 3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, search(tc.tree.root, tc.key))
		})
	}
}

func TestTree_min(t *testing.T) {
	tt := []struct {
		name string
		tree *BSTree
		want *node
	}{
		{
			name: "Nil node returns nil",
			tree: &BSTree{},
		},
		{
			name: "root returns root value",
			tree: &BSTree{
				root: &node{
					key: 5,
				},
			},
			want: &node{key: 5},
		},
		{
			name: "h=1 returns correct value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 15,
					},
				},
			},
			want: &node{key: 5},
		},
		{
			name: "h=2 returns correct value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
						left: &node{
							key: 3,
						},
						right: &node{
							key: 7,
						},
					},
					right: &node{
						key: 15,
						left: &node{
							key: 12,
						},
						right: &node{
							key: 19,
						},
					},
				},
			},
			want: &node{key: 3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, min(tc.tree.root))
		})
	}
}

func TestTree_max(t *testing.T) {
	tt := []struct {
		name string
		tree *BSTree
		want *node
	}{
		{
			name: "Nil node returns nil",
			tree: &BSTree{},
		},
		{
			name: "root returns root value",
			tree: &BSTree{
				root: &node{
					key: 5,
				},
			},
			want: &node{key: 5},
		},
		{
			name: "h=1 returns correct value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 15,
					},
				},
			},
			want: &node{key: 15},
		},
		{
			name: "h=2 returns correct value",
			tree: &BSTree{
				root: &node{
					key: 10,
					left: &node{
						key: 5,
						left: &node{
							key: 3,
						},
						right: &node{
							key: 7,
						},
					},
					right: &node{
						key: 15,
						left: &node{
							key: 12,
						},
						right: &node{
							key: 19,
						},
					},
				},
			},
			want: &node{key: 19},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, max(tc.tree.root))
		})
	}
}

func TestTree_Upsert(t *testing.T) {
	t.Run("insert", func(t *testing.T) {
		tt := []struct {
			name  string
			items []item
		}{
			{
				name: "insert one",
				items: []item{
					{
						key:     5,
						payload: "test",
					},
				},
			},
			{
				name: "insert two",
				items: []item{
					{
						key:     5,
						payload: "test",
					},
					{
						key:     10,
						payload: "test2",
					},
				},
			},
			{
				name: "insert four",
				items: []item{
					{
						key:     5,
						payload: "test",
					},
					{
						key:     10,
						payload: "test2",
					},
					{
						key:     15,
						payload: "test3",
					},
					{
						key:     3,
						payload: "test4",
					},
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				tree := NewBSTree()

				for _, i := range tc.items {
					tree.Upsert(i.key, i.payload)
				}

				for i, v := range tc.items {
					assert.Equal(t, tc.items[i].payload, tree.Search(v.key))
				}
			})
		}
	})

	t.Run("update", func(t *testing.T) {
		tree := NewBSTree()

		tree.Upsert(1, "test")
		tree.Upsert(2, "test2")

		assert.Equal(t, "test", tree.Search(1))
		assert.Equal(t, "test2", tree.Search(2))

		tree.Upsert(1, "test3")

		assert.Equal(t, "test3", tree.Search(1))
		assert.Equal(t, "test2", tree.Search(2))
	})
}

func TestTree_Height(t *testing.T) {
	tt := []struct {
		name  string
		items []int64
		want  int
	}{
		{
			name: "empty tree returns -1",
			want: -1,
		},
		{
			name:  "rooted tree returns 0",
			items: []int64{15},
			want:  0,
		},
		{
			name:  "root with two children returns 1",
			items: []int64{15, 10, 20},
			want:  1,
		},
		{
			name:  "left-sided tree returns 2",
			items: []int64{15, 10, 5},
			want:  2,
		},
		{
			name:  "right-sided tree returns 2",
			items: []int64{15, 20, 25},
			want:  2,
		},
		{
			name:  "custom tree returns 4",
			items: []int64{15, 6, 18, 3, 7, 17, 20, 2, 4, 13, 19, 9},
			want:  4,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBSTree()

			for _, v := range tc.items {
				tree.Upsert(v, nil)
			}

			assert.Equal(t, tc.want, tree.Height())
		})
	}
}

func TestTree_Successor(t *testing.T) {
	tt := []struct {
		name    string
		items   []item
		toCheck int64
		want    string
	}{
		{
			name: "empty tree yields nil",
		},
		{
			name: "h=0 tree yields nil",
			items: []item{
				{
					key:     15,
					payload: "test",
				},
			},
			toCheck: 15,
		},
		{
			name: "h=1 tree yields result",
			items: []item{
				{
					key:     15,
					payload: "test",
				},
				{
					key:     20,
					payload: "test2",
				},
				{
					key:     5,
					payload: "test3",
				},
			},
			toCheck: 5,
			want:    "test",
		},
		{
			name:    "multi-level tree yields result",
			toCheck: 4,
			want:    "6",
			items: []item{
				{
					key:     15,
					payload: "15",
				},
				{
					key:     6,
					payload: "6",
				},
				{
					key:     18,
					payload: "18",
				},
				{
					key:     3,
					payload: "3",
				},
				{
					key:     7,
					payload: "7",
				},
				{
					key:     17,
					payload: "17",
				},
				{
					key:     20,
					payload: "20",
				},
				{
					key:     2,
					payload: "2",
				},
				{
					key:     4,
					payload: "4",
				},
				{
					key:     13,
					payload: "13",
				},
				{
					key:     19,
					payload: "19",
				},
				{
					key:     9,
					payload: "9",
				},
			},
		},
		{
			name:    "multi-level tree yields result",
			toCheck: 13,
			want:    "15",
			items: []item{
				{
					key:     15,
					payload: "15",
				},
				{
					key:     6,
					payload: "6",
				},
				{
					key:     18,
					payload: "18",
				},
				{
					key:     3,
					payload: "3",
				},
				{
					key:     7,
					payload: "7",
				},
				{
					key:     17,
					payload: "17",
				},
				{
					key:     20,
					payload: "20",
				},
				{
					key:     2,
					payload: "2",
				},
				{
					key:     4,
					payload: "4",
				},
				{
					key:     13,
					payload: "13",
				},
				{
					key:     19,
					payload: "19",
				},
				{
					key:     9,
					payload: "9",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBSTree()

			for _, i := range tc.items {
				tree.Upsert(i.key, i.payload)
			}

			if tc.want != "" {
				assert.Equal(t, tc.want, tree.Successor(tc.toCheck))
			} else {
				assert.Nil(t, tree.Successor(tc.toCheck))
			}
		})
	}
}

func TestTree_Delete(t *testing.T) {
	t.Run("delete on empty tree is noop", func(t *testing.T) {
		tree := NewBSTree()

		assert.Equal(t, -1, tree.Height())

		tree.Delete(0)

		assert.Equal(t, -1, tree.Height())
	})

	t.Run("delete on rooted tree on non-existing key is noop", func(t *testing.T) {
		tree := NewBSTree()

		tree.Upsert(15, "test")
		tree.Delete(0)

		assert.Equal(t, 0, tree.Height())
		assert.Equal(t, "test", tree.Search(15))
	})

	t.Run("deleting root node leaves empty tree", func(t *testing.T) {
		tree := NewBSTree()

		tree.Upsert(15, "test")
		assert.Equal(t, 0, tree.Height())

		tree.Delete(15)

		assert.Equal(t, -1, tree.Height())
		assert.Nil(t, tree.root)
	})

	t.Run("deleted node gets replace by right child", func(t *testing.T) {
		tree := NewBSTree()

		z := &node{key: 15}
		r := &node{key: 20}
		r1 := &node{key: 25}
		r2 := &node{key: 23}

		tree.root = z
		z.right = r

		r.parent = z
		r.right = r1
		r.left = r2

		assert.Equal(t, 2, tree.Height())

		tree.Delete(15)

		// r is now root
		assert.Nil(t, r.parent)
		assert.Equal(t, 1, tree.Height())
		assert.Equal(t, r1, r.right)
		assert.Equal(t, r2, r.left)
	})

	t.Run("deleted node gets replace by left child", func(t *testing.T) {
		tree := NewBSTree()

		z := &node{key: 15}
		r := &node{key: 5}
		r1 := &node{key: 10}
		r2 := &node{key: 3}

		tree.root = z
		z.left = r

		r.parent = z
		r.right = r1
		r.left = r2

		assert.Equal(t, 2, tree.Height())

		tree.Delete(15)

		// r is now root
		assert.Nil(t, r.parent)
		assert.Equal(t, 1, tree.Height())
		assert.Equal(t, r1, r.right)
		assert.Equal(t, r2, r.left)
	})

	t.Run("deleted nodes has two children, successor is right child", func(t *testing.T) {
		tree := NewBSTree()

		z := &node{key: 15}
		y := &node{key: 20}
		l := &node{key: 5}
		x := &node{key: 25}

		tree.root = z
		z.left = l
		z.right = y

		l.parent = z

		y.parent = z
		y.right = x

		assert.Equal(t, 2, tree.Height())

		tree.Delete(15)

		// y is now parent with l as its left subtree
		assert.Nil(t, y.parent)
		assert.Equal(t, 1, tree.Height())
		assert.Equal(t, l, y.left)
		assert.Equal(t, x, y.right)
	})

	t.Run("deleted node has two children, successor in the left subtree of root's right child", func(t *testing.T) {
		tree := NewBSTree()

		z := &node{key: 15} // To delete
		l := &node{key: 5}
		r := &node{key: 30} // will be right child of successor
		u := &node{key: 40}
		y := &node{key: 25} // successor
		x := &node{key: 28} // new left subtree of r

		tree.root = z
		z.left = l
		z.right = r

		l.parent = z

		r.parent = z
		r.right = u
		r.left = y

		y.parent = r
		y.right = x

		assert.Equal(t, 3, tree.Height())

		tree.Delete(15)

		assert.Nil(t, y.parent)
		assert.Equal(t, 2, tree.Height())
		assert.Equal(t, r, y.right)
		assert.Equal(t, x, r.left)
	})
}
