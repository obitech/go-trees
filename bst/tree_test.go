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
			key:  1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.tree.Search(tc.key))
		})
	}
}
