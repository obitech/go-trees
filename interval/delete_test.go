package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTree_Delete(t *testing.T) {
	t.Run("delete on empty tree is noop", func(t *testing.T) {
		tree := NewIntervalTree()

		assert.Equal(t, -1, tree.Height())

		tree.Delete(Interval{})

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("deleting unknown node is noop", func(t *testing.T) {
		tree := NewIntervalTree()

		root, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		rootL, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		rootR, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))

		tree.Upsert(root, "Aug")
		tree.Upsert(rootL, "Mar")
		tree.Upsert(rootR, "Sep")

		assert.Equal(t, 1, tree.Height())

		del, _ := NewInterval(newTime(t, "2019-Sep-01"), newTime(t, "2019-Sep-02"))
		tree.Delete(del)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, root, tree.root.key)
		assert.Equal(t, rootL, tree.root.left.key)
		assert.Equal(t, rootR, tree.root.right.key)
	})

	t.Run("deleting root node leaves empty tree", func(t *testing.T) {
		tree := NewIntervalTree()

		root, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		tree.Upsert(root, "Aug")

		assert.Equal(t, 0, tree.Height())

		tree.Delete(root)

		assert.Equal(t, -1, tree.Height())
		assert.Equal(t, tree.sentinel, tree.root)
	})

	t.Run("(h=1) delete right leaf is successful", func(t *testing.T) {
		tree := NewIntervalTree()

		root, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		rootL, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		rootR, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))

		tree.Upsert(root, "Aug")
		tree.Upsert(rootL, "Mar")
		tree.Upsert(rootR, "Sep")

		tree.Delete(rootR)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, root, tree.root.key)
		assert.Equal(t, rootL, tree.root.left.key)
		assert.Equal(t, tree.sentinel, tree.root.right)
	})

	t.Run("(h=1) delete left leaf is successful", func(t *testing.T) {
		tree := NewIntervalTree()

		root, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		rootL, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		rootR, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))

		tree.Upsert(root, "Aug")
		tree.Upsert(rootL, "Mar")
		tree.Upsert(rootR, "Sep")

		tree.Delete(rootL)

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, root, tree.root.key)
		assert.Equal(t, rootR, tree.root.right.key)
		assert.Equal(t, tree.sentinel, tree.root.left)
	})

	t.Run("(h=3) Delete from tree is successful", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		feb, _ := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))
		dec, _ := NewInterval(newTime(t, "2020-Dec-01"), newTime(t, "2020-Dec-02"))
		mar, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		jan, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))
		jul, _ := NewInterval(newTime(t, "2020-Jul-01"), newTime(t, "2020-Jul-02"))
		may, _ := NewInterval(newTime(t, "2020-May-01"), newTime(t, "2020-May-02"))
		aug, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		apr, _ := NewInterval(newTime(t, "2020-Apr-01"), newTime(t, "2020-Apr-02"))
		jun, _ := NewInterval(newTime(t, "2020-Jun-01"), newTime(t, "2020-Jun-02"))
		oct, _ := NewInterval(newTime(t, "2020-Oct-01"), newTime(t, "2020-Oct-02"))

		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")
		tree.Upsert(dec, "Dec")
		tree.Upsert(mar, "Mar")
		tree.Upsert(jan, "Jan")
		tree.Upsert(jul, "Jul")
		tree.Upsert(may, "May")
		tree.Upsert(aug, "Aug")
		tree.Upsert(apr, "Apr")
		tree.Upsert(jun, "Jun")
		tree.Upsert(oct, "Oct")

		tree.Delete(jul)
		tree.Delete(nov)
		tree.Delete(dec)
		tree.Delete(mar)
		tree.Delete(feb)

		root := tree.root               // May
		rootL := tree.root.left         // Apr
		rootLL := tree.root.left.left   // Jan
		rootR := tree.root.right        // Aug
		rootRL := tree.root.right.left  // Jun
		rootRR := tree.root.right.right // Oct

		assert.Equal(t, 2, tree.Height())

		assert.Equal(t, tree.sentinel, root.parent)
		assert.Equal(t, black, root.color)
		assert.Equal(t, may, root.key)
		assert.Equal(t, oct.high, root.max)
		assert.Equal(t, "May", root.payload)

		assert.Equal(t, root, rootL.parent)
		assert.Equal(t, black, rootL.color)
		assert.Equal(t, apr, rootL.key)
		assert.Equal(t, apr.high, rootL.max)
		assert.Equal(t, "Apr", rootL.payload)

		assert.Equal(t, rootL, rootLL.parent)
		assert.Equal(t, red, rootLL.color)
		assert.Equal(t, jan, rootLL.key)
		assert.Equal(t, jan.high, rootLL.max)
		assert.Equal(t, "Jan", rootLL.payload)

		assert.Equal(t, root, rootR.parent)
		assert.Equal(t, red, rootR.color)
		assert.Equal(t, aug, rootR.key)
		assert.Equal(t, oct.high, rootR.max)
		assert.Equal(t, "Aug", rootR.payload)

		assert.Equal(t, rootR, rootRL.parent)
		assert.Equal(t, black, rootRL.color)
		assert.Equal(t, jun, rootRL.key)
		assert.Equal(t, jun.high, rootRL.max)
		assert.Equal(t, "Jun", rootRL.payload)

		assert.Equal(t, rootR, rootRR.parent)
		assert.Equal(t, black, rootRR.color)
		assert.Equal(t, oct, rootRR.key)
		assert.Equal(t, oct.high, rootRR.max)
		assert.Equal(t, "Oct", rootRR.payload)
	})
}
