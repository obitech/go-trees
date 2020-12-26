package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Search(t *testing.T) {
	tt := []struct {
		name string
		tree *Tree
		key  int64
		want *Node
	}{
		{
			name: "nil tree returns nil",
			tree: &Tree{},
		},
		{
			name: "root hit returns value",
			tree: &Tree{
				Root: &Node{
					Key: 5,
				},
			},
			key:  5,
			want: &Node{Key: 5},
		},
		{
			name: "root miss returns nil",
			tree: &Tree{
				Root: &Node{
					Key: 6,
				},
			},
			key: 5,
		},
		{
			name: "h=1 left tree hit returns value",
			tree: &Tree{
				Root: &Node{
					Key: 10,
					Left: &Node{
						Key: 5,
					},
					Right: &Node{
						Key: 15,
					},
				},
			},
			key:  5,
			want: &Node{Key: 5},
		},
		{
			name: "h=2 right tree hit returns value",
			tree: &Tree{
				Root: &Node{
					Key: 10,
					Left: &Node{
						Key: 5,
					},
					Right: &Node{
						Key: 15,
					},
				},
			},
			key:  15,
			want: &Node{Key: 15},
		},
		{
			name: "h=2 miss returns nil",
			tree: &Tree{
				Root: &Node{
					Key: 10,
					Left: &Node{
						Key: 5,
					},
					Right: &Node{
						Key: 15,
					},
				},
			},
			key: 99,
		},
		{
			name: "h=3 right tree hit returns value",
			tree: &Tree{
				Root: &Node{
					Key: 10,
					Left: &Node{
						Key: 5,
						Left: &Node{
							Key: 3,
						},
						Right: &Node{
							Key: 7,
						},
					},
					Right: &Node{
						Key: 15,
						Left: &Node{
							Key: 12,
						},
						Right: &Node{
							Key: 19,
						},
					},
				},
			},
			key:  19,
			want: &Node{Key: 19},
		},
		{
			name: "h=3 left tree hit returns value",
			tree: &Tree{
				Root: &Node{
					Key: 10,
					Left: &Node{
						Key: 5,
						Left: &Node{
							Key: 3,
						},
						Right: &Node{
							Key: 7,
						},
					},
					Right: &Node{
						Key: 15,
						Left: &Node{
							Key: 12,
						},
						Right: &Node{
							Key: 19,
						},
					},
				},
			},
			key:  3,
			want: &Node{Key: 3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.tree.Search(tc.key))
		})
	}
}

func TestTree_recursiveAndIterativeSearch(t *testing.T) {
	tt := []struct {
		name string
		node *Node
		key  int64
		want *Node
	}{
		{
			name: "nil node returns nil",
		},
		{
			name: "root hit returns value",
			node: &Node{
				Key: 5,
			},
			key:  5,
			want: &Node{Key: 5},
		},
		{
			name: "root miss returns nil",
			node: &Node{
				Key: 6,
			},
			key: 5,
		},
		{
			name: "h=1 left tree hit returns value",
			node: &Node{
				Key: 10,
				Left: &Node{
					Key: 5,
				},
				Right: &Node{
					Key: 1,
				},
			},
			key:  5,
			want: &Node{Key: 5},
		},
		{
			name: "h=2 right tree hit returns value",
			node: &Node{
				Key: 10,
				Left: &Node{
					Key: 5,
				},
				Right: &Node{
					Key: 15,
				},
			},
			key:  15,
			want: &Node{Key: 15},
		},
		{
			name: "h=2 miss returns nil",
			node: &Node{
				Key: 10,
				Left: &Node{
					Key: 5,
				},
				Right: &Node{
					Key: 15,
				},
			},
			key: 99,
		},
		{
			name: "h=3 right tree hit returns value",
			node: &Node{
				Key: 10,
				Left: &Node{
					Key: 5,
					Left: &Node{
						Key: 3,
					},
					Right: &Node{
						Key: 7,
					},
				},
				Right: &Node{
					Key: 15,
					Left: &Node{
						Key: 12,
					},
					Right: &Node{
						Key: 19,
					},
				},
			},
			key:  19,
			want: &Node{Key: 19},
		},
		{
			name: "h=3 left tree hit returns value",
			node: &Node{
				Key: 10,
				Left: &Node{
					Key: 5,
					Left: &Node{
						Key: 3,
					},
					Right: &Node{
						Key: 7,
					},
				},
				Right: &Node{
					Key: 15,
					Left: &Node{
						Key: 12,
					},
					Right: &Node{
						Key: 19,
					},
				},
			},
			key:  3,
			want: &Node{Key: 3},
		},
	}

	for _, tc := range tt {
		t.Run("searchRecursive", func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.want, searchRecursive(tc.node, tc.key))
			})
		})

		t.Run("searchIterative", func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.want, searchIterative(tc.node, tc.key))
			})
		})
	}
}
