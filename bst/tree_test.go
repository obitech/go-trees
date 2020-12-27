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
		tree *Tree
		key  int64
		want *node
	}{
		{
			name: "nil tree returns nil",
			tree: &Tree{},
		},
		{
			name: "root hit returns value",
			tree: &Tree{
				root: &node{
					key: 5,
				},
			},
			key:  5,
			want: &node{key: 5},
		},
		{
			name: "root miss returns nil",
			tree: &Tree{
				root: &node{
					key: 6,
				},
			},
			key: 5,
		},
		{
			name: "h=1 left tree hit returns value",
			tree: &Tree{
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
			tree: &Tree{
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
			tree: &Tree{
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
			tree: &Tree{
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
			tree: &Tree{
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
		tree *Tree
		want *node
	}{
		{
			name: "Nil node returns nil",
			tree: &Tree{},
		},
		{
			name: "root returns root value",
			tree: &Tree{
				root: &node{
					key: 5,
				},
			},
			want: &node{key: 5},
		},
		{
			name: "h=1 returns correct value",
			tree: &Tree{
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
			tree: &Tree{
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
		tree *Tree
		want *node
	}{
		{
			name: "Nil node returns nil",
			tree: &Tree{},
		},
		{
			name: "root returns root value",
			tree: &Tree{
				root: &node{
					key: 5,
				},
			},
			want: &node{key: 5},
		},
		{
			name: "h=1 returns correct value",
			tree: &Tree{
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
			tree: &Tree{
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
				tree := NewBST()

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
		tree := NewBST()

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
			tree := NewBST()

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
			tree := NewBST()

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
