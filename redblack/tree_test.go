package redblack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Root(t *testing.T) {
	t.Run("empty tree returns nil as payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		assert.Nil(t, tree.Root())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("rooted tree returns root payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.root = tree.newLeaf(myInt(15), "root")
		tree.root.parent = tree.sentinel

		assert.Equal(t, "root", tree.Root())
	})
}

func TestTree_Height(t *testing.T) {
	t.Run("empty tree returns -1", func(t *testing.T) {
		tree := NewRedBlackTree()
		assert.Equal(t, -1, tree.Height())
	})

	t.Run("rooted tree returns 0", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), nil)
		x.parent = tree.sentinel

		tree.root = x

		assert.Equal(t, 0, tree.Height())
	})

	t.Run("h=1 tree returns correct height", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), nil)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), nil)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), nil)
		r.parent = x
		x.right = r

		assert.Equal(t, 1, tree.Height())
	})

	t.Run("h=2 tree returns correct height", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), nil)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), nil)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), nil)
		r.parent = x
		x.right = r

		y := tree.newLeaf(myInt(25), nil)
		y.parent = r
		r.right = y

		z := tree.newLeaf(myInt(9), nil)
		z.parent = l
		l.left = z

		assert.Equal(t, 2, tree.Height())
	})
}

func TestTree_Min(t *testing.T) {
	t.Run("empty tree returns nil", func(t *testing.T) {
		tree := NewRedBlackTree()

		assert.Nil(t, tree.Min())
	})

	t.Run("rooted tree returns root payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.root = tree.newLeaf(myInt(15), "root")
		tree.root.parent = tree.sentinel

		assert.Equal(t, "root", tree.Min())
	})

	t.Run("h=1 tree returns correct Min payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), 15)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), 12)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), 20)
		r.parent = x
		x.right = r

		assert.Equal(t, 12, tree.Min())
	})

	t.Run("h=2 tree returns correct Min payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), 15)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), 12)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), 20)
		r.parent = x
		x.right = r

		y := tree.newLeaf(myInt(25), 25)
		y.parent = r
		r.right = y

		z := tree.newLeaf(myInt(9), 9)
		z.parent = l
		l.left = z

		assert.Equal(t, 9, tree.Min())
	})
}

func TestTree_Max(t *testing.T) {
	t.Run("empty tree returns nil", func(t *testing.T) {
		tree := NewRedBlackTree()

		assert.Nil(t, tree.Max())
	})

	t.Run("rooted tree returns root payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		tree.root = tree.newLeaf(myInt(15), "root")
		tree.root.parent = tree.sentinel

		assert.Equal(t, "root", tree.Max())
	})

	t.Run("h=1 tree returns correct Max payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), 15)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), 12)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), 20)
		r.parent = x
		x.right = r

		assert.Equal(t, 20, tree.Max())
	})

	t.Run("h=2 tree returns correct Max payload", func(t *testing.T) {
		tree := NewRedBlackTree()

		x := tree.newLeaf(myInt(15), 15)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(myInt(12), 12)
		l.parent = x
		x.left = l

		r := tree.newLeaf(myInt(20), 20)
		r.parent = x
		x.right = r

		y := tree.newLeaf(myInt(25), 25)
		y.parent = r
		r.right = y

		z := tree.newLeaf(myInt(9), 9)
		z.parent = l
		l.left = z

		assert.Equal(t, 25, tree.Max())
	})
}

func TestTree_rotateLeft(t *testing.T) {
	tree := NewRedBlackTree()

	x := tree.newLeaf(myInt(15), 15)
	y := tree.newLeaf(myInt(30), 30)

	xL := tree.newLeaf(myInt(10), 10)
	xL.parent = x
	x.left = xL
	x.right = y

	yL := tree.newLeaf(myInt(25), 25)
	yL.parent = y

	yR := tree.newLeaf(myInt(35), 35)
	yR.parent = y
	y.left = yL
	y.right = yR

	tree.root = x
	x.parent = tree.sentinel

	assert.Equal(t, 2, tree.Height())

	tree.rotateLeft(x)

	assert.Equal(t, 2, tree.Height())
	assert.Equal(t, tree.root, y)
	assert.Equal(t, y.left, x)
	assert.Equal(t, y.right, yR)
	assert.Equal(t, x.left, xL)
	assert.Equal(t, x.right, yL)
}

func TestTree_rotateRight(t *testing.T) {
	tree := NewRedBlackTree()

	assert.Equal(t, -1, tree.Height())

	x := tree.newLeaf(myInt(20), 20)
	y := tree.newLeaf(myInt(40), 40)

	xR := tree.newLeaf(myInt(25), 25)
	xR.parent = x
	x.left = y
	x.right = xR

	yR := tree.newLeaf(myInt(50), 50)
	yR.parent = y

	yL := tree.newLeaf(myInt(45), 45)
	yL.parent = y

	y.left = yL
	y.right = yR

	tree.root = x
	x.parent = tree.sentinel

	assert.Equal(t, 2, tree.Height())

	tree.rotateRight(x)

	assert.Equal(t, 2, tree.Height())
	assert.Equal(t, tree.root, y)
	assert.Equal(t, y.left, yL)
	assert.Equal(t, y.right, x)
	assert.Equal(t, x.left, yR)
	assert.Equal(t, x.right, xR)
}

func TestTree_Successor(t *testing.T) {
	tree := NewRedBlackTree()
	for _, i := range []int64{1, 20, 3, 5, 21, 12, 18, 13, 4, 8, 50, 30} {
		tree.Upsert(myInt(i), i)
	}

	tt := []struct {
		name  string
		check int64
		want  int64
	}{
		{
			name:  "1 -> 3",
			check: 1,
			want:  3,
		},
		{
			name:  "3 -> 4",
			check: 3,
			want:  4,
		},
		{
			name:  "4 -> 5",
			check: 4,
			want:  5,
		},
		{
			name:  "5 -> 8",
			check: 5,
			want:  8,
		},
		{
			name:  "8 -> 12",
			check: 8,
			want:  12,
		},
		{
			name:  "12 -> 13",
			check: 12,
			want:  13,
		},
		{
			name:  "13 -> 18",
			check: 13,
			want:  18,
		},
		{
			name:  "18 -> 20",
			check: 18,
			want:  20,
		},
		{
			name:  "20 -> 21",
			check: 20,
			want:  21,
		},
		{
			name:  "21 -> 30",
			check: 21,
			want:  30,
		},
		{
			name:  "30 -> 50",
			check: 30,
			want:  50,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tree.Successor(myInt(tc.check)))
		})
	}
}
