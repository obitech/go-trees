package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTree_updateMax(t *testing.T) {
	t.Run("Root node without children doesn't get updated", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))

		tree.Upsert(nov, "Nov")

		tree.updateMax(tree.root)

		assert.Equal(t, tree.root.max, nov.high)
	})

	t.Run("node has left child which is higher, updates node max", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		feb, _ := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))

		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")

		n := newTime(t, "2021-Feb-01")

		tree.root.left.max = n

		tree.updateMax(tree.root)

		assert.Equal(t, n, tree.root.max)
	})

	t.Run("node has right child which is higher, updates node max", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		feb, _ := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))
		oct, _ := NewInterval(newTime(t, "2020-Oct-01"), newTime(t, "2020-Oct-02"))

		tree.Upsert(oct, "Oct")
		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")

		n, _ := NewInterval(newTime(t, "2019-Nov-01"), newTime(t, "2019-Nov-02"))

		tree.root.key = n
		tree.root.max = n.high

		tree.updateMax(tree.root)

		assert.Equal(t, nov.high, tree.root.max)
	})

	t.Run("node max is higher, updates node max", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		feb, _ := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))

		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")

		n, _ := NewInterval(newTime(t, "2019-Nov-01"), newTime(t, "2019-Nov-02"))

		tree.root.key = n
		tree.root.max = n.high

		tree.updateMax(tree.root)

		assert.Equal(t, feb.high, tree.root.max)
	})
}

func TestIntervalTree_Root(t *testing.T) {
	t.Run("empty tree returns nil as payload", func(t *testing.T) {
		tree := NewIntervalTree()

		assert.Equal(t, tree.sentinel, tree.root)

		assert.Equal(t, -1, tree.Height())

		v, err := tree.Root()
		assert.Equal(t, Result{}, v)
		assert.IsType(t, ErrNotFound(""), err)
	})

	t.Run("rooted tree returns root payload", func(t *testing.T) {
		tree := NewIntervalTree()

		i, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))

		tree.root = tree.newLeaf(i, "root")
		tree.root.parent = tree.sentinel

		assert.Equal(t, 0, tree.Height())

		v, err := tree.Root()
		assert.NoError(t, err)
		assert.Equal(t, Result{
			Interval: i,
			Payload:  "root",
		}, v)

		t.Run("modifying Result does not modify tree", func(t *testing.T) {
			v.Payload = "foo"

			check, err := tree.Root()
			assert.NoError(t, err)
			assert.Equal(t, "root", check.Payload)
		})
	})
}

func TestIntervalTree_Height(t *testing.T) {
	t.Run("empty tree returns -1", func(t *testing.T) {
		tree := NewIntervalTree()
		assert.Equal(t, -1, tree.Height())
	})

	t.Run("rooted tree returns 0", func(t *testing.T) {
		tree := NewIntervalTree()

		x := tree.newLeaf(Interval{}, nil)
		x.parent = tree.sentinel

		tree.root = x

		assert.Equal(t, 0, tree.Height())
	})

	t.Run("h=1 tree returns correct height", func(t *testing.T) {
		tree := NewIntervalTree()

		x := tree.newLeaf(Interval{}, nil)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(Interval{}, nil)
		l.parent = x
		x.left = l

		r := tree.newLeaf(Interval{}, nil)
		r.parent = x
		x.right = r

		assert.Equal(t, 1, tree.Height())
	})

	t.Run("h=2 tree returns correct height", func(t *testing.T) {
		tree := NewIntervalTree()

		x := tree.newLeaf(Interval{}, nil)
		x.parent = tree.sentinel
		tree.root = x

		l := tree.newLeaf(Interval{}, nil)
		l.parent = x
		x.left = l

		r := tree.newLeaf(Interval{}, nil)
		r.parent = x
		x.right = r

		y := tree.newLeaf(Interval{}, nil)
		y.parent = r
		r.right = y

		z := tree.newLeaf(Interval{}, nil)
		z.parent = l
		l.left = z

		assert.Equal(t, 2, tree.Height())
	})
}

func TestIntervalTree_rotateLeft(t *testing.T) {
	tree := NewIntervalTree()

	iX, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
	iXL, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))

	iY, _ := NewInterval(newTime(t, "2020-Oct-01"), newTime(t, "2020-Oct-02"))
	iYL, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))
	iYR, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))

	x := tree.newLeaf(iX, 15)
	x.max = newTime(t, "2020-Aug-02")

	y := tree.newLeaf(iY, 30)
	y.max = newTime(t, "2020-Nov-02")

	xL := tree.newLeaf(iXL, 10)
	xL.parent = x
	xL.max = newTime(t, "2020-Jan-02")

	x.left = xL
	x.right = y

	yL := tree.newLeaf(iYL, 25)
	yL.parent = y
	yL.max = newTime(t, "2020-Sep-02")

	yR := tree.newLeaf(iYR, 35)
	yR.parent = y
	yR.max = newTime(t, "2020-Nov-02")

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

	assert.Equal(t, newTime(t, "2020-Sep-02"), x.max)
	assert.Equal(t, newTime(t, "2020-Nov-02"), y.max)
}

func TestIntervalTree_rotateRight(t *testing.T) {
	tree := NewIntervalTree()

	assert.Equal(t, -1, tree.Height())

	iX, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
	iXR, _ := NewInterval(newTime(t, "2020-Oct-01"), newTime(t, "2020-Oct-02"))

	iY, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
	iYL, _ := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))
	iYR, _ := NewInterval(newTime(t, "2020-Apr-01"), newTime(t, "2020-Apr-02"))

	x := tree.newLeaf(iX, "Aug")
	x.max = newTime(t, "2020-Oct-02")

	y := tree.newLeaf(iY, "Mar")
	y.max = newTime(t, "2020-Apr-02")

	xR := tree.newLeaf(iXR, "Oct")
	xR.parent = x
	xR.max = newTime(t, "2020-Oct-02")

	x.left = y
	x.right = xR

	yR := tree.newLeaf(iYR, "Apr")
	yR.parent = y
	yR.max = newTime(t, "2020-Apr-02")

	yL := tree.newLeaf(iYL, "Feb")
	yL.parent = y
	yL.max = newTime(t, "2020-Feb-02")

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

	assert.Equal(t, newTime(t, "2020-Oct-02"), y.max)
	assert.Equal(t, newTime(t, "2020-Oct-02"), x.max)
}

func TestIntervalTree_Min(t *testing.T) {
	// TODO: write tests
}
