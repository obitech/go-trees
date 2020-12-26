package interval

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTree_Upsert(t *testing.T) {
	t.Run("Insert into empty tree creates root node", func(t *testing.T) {
		tree := NewIntervalTree()

		assert.Equal(t, -1, tree.Height())

		i, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))

		tree.Upsert(i, "root")

		assert.Equal(t, 0, tree.Height())
		assert.Equal(t, i, tree.root.key)
		assert.Equal(t, newTime(t, "2020-Jan-02"), tree.root.max)
		assert.Equal(t, tree.sentinel, tree.root.left)
		assert.Equal(t, tree.sentinel, tree.root.right)

		t.Run("searching for root node succeeds", func(t *testing.T) {
			v, err := tree.FindFirstOverlapping(i)

			assert.NoError(t, err)
			assert.Equal(t, Result{Interval: i, Payload: "root"}, v)
		})

		t.Run("searching for non-overlapping interval yields nil", func(t *testing.T) {
			v, _ := NewInterval(newTime(t, "2019-Jan-01"), newTime(t, "2019-Feb-01"))

			r, err := tree.FindFirstOverlapping(v)
			assert.Equal(t, r, Result{})
			assert.IsType(t, ErrNotFound(""), err)
		})
	})
	t.Run("exact interval upsert updates root node", func(t *testing.T) {
		tree := NewIntervalTree()

		i, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))
		tree.Upsert(i, "root")
		tree.Upsert(i, "root2")

		assert.Equal(t, 0, tree.Height())

		r, err := tree.Root()
		assert.Nil(t, err)
		assert.Equal(t, "root2", r.Payload)
	})

	t.Run("overlapping, non-exact interval does not update root node", func(t *testing.T) {
		tree := NewIntervalTree()

		i, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-02"))
		tree.Upsert(i, "root")

		i2, _ := NewInterval(newTime(t, "2020-Jan-01"), newTime(t, "2020-Jan-03"))
		tree.Upsert(i2, "root2")

		assert.Equal(t, 1, tree.Height())

		r, err := tree.Root()
		assert.Nil(t, err)
		assert.Equal(t, "root", r.Payload)

		assert.Equal(t, newTime(t, "2020-Jan-03"), tree.root.max)

		assert.Equal(t, newTime(t, "2020-Jan-03"), tree.root.right.max)
		assert.Equal(t, newTime(t, "2020-Jan-03"), tree.root.right.key.high)
	})

	t.Run("(h=1) Insert succeeds", func(t *testing.T) {
		tree := NewIntervalTree()

		x, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		xL, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		xR, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))

		tree.Upsert(x, "Aug")
		tree.Upsert(xL, "Mar")
		tree.Upsert(xR, "Sep")

		assert.Equal(t, 1, tree.Height())

		assert.Equal(t, black, tree.root.color, "root is black")
		assert.Equal(t, "Aug", tree.root.payload, "root is Aug")
		assert.Equal(t, newTime(t, "2020-Sep-02"), tree.root.max, "root.max is Sep")
		assert.Equal(t, tree.sentinel, tree.root.parent)

		assert.Equal(t, red, tree.root.left.color, "left node is red")
		assert.Equal(t, "Mar", tree.root.left.payload, "left node is Mar")
		assert.Equal(t, newTime(t, "2020-Mar-02"), tree.root.left.max, "left.max is Mar")
		assert.Equal(t, tree.root.left.parent, tree.root)

		assert.Equal(t, red, tree.root.right.color, "right node is red")
		assert.Equal(t, "Sep", tree.root.right.payload, "right node is Sep")
		assert.Equal(t, newTime(t, "2020-Sep-02"), tree.root.right.max, "right.max is Sep")
		assert.Equal(t, tree.root.right.parent, tree.root)

		t.Run("searching overlapping interval yields result", func(t *testing.T) {
			i, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-05"))

			v, err := tree.FindFirstOverlapping(i)
			assert.NoError(t, err)
			assert.Equal(t, Result{
				Interval: xR,
				Payload:  "Sep",
			}, v)
		})
	})

	t.Run("(h=2) Insert with same lower bound", func(t *testing.T) {
		tree := NewIntervalTree()

		nov1, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		nov2, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-03"))

		tree.Upsert(nov1, "Nov1")
		tree.Upsert(nov2, "Nov2")

		assert.Equal(t, 1, tree.Height())
	})

	t.Run("(h=3) Inserting into tree is successful", func(t *testing.T) {
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

		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")
		tree.Upsert(dec, "Dec")
		tree.Upsert(mar, "Mar")
		tree.Upsert(jan, "Jan")
		tree.Upsert(jul, "Jul")
		tree.Upsert(may, "May")
		tree.Upsert(aug, "Aug")
		tree.Upsert(apr, "Apr")

		root := tree.root                     // May
		rootL := tree.root.left               // Feb
		rootR := tree.root.right              // Nov
		rootLL := tree.root.left.left         // Jan
		rootLR := tree.root.left.right        // Mar
		rootLRR := tree.root.left.right.right // Apr
		rootRL := tree.root.right.left        // Jul
		rootRLR := tree.root.right.left.right // Aug
		rootRR := tree.root.right.right       // Dec

		assert.Equal(t, 3, tree.Height())

		assert.Equal(t, tree.sentinel, root.parent)
		assert.Equal(t, black, root.color)
		assert.Equal(t, may, root.key)
		assert.Equal(t, dec.high, root.max)
		assert.Equal(t, "May", root.payload)

		assert.Equal(t, root, rootL.parent)
		assert.Equal(t, red, rootL.color)
		assert.Equal(t, feb, rootL.key)
		assert.Equal(t, apr.high, rootL.max)
		assert.Equal(t, "Feb", rootL.payload)

		assert.Equal(t, rootL, rootLL.parent)
		assert.Equal(t, black, rootLL.color)
		assert.Equal(t, jan, rootLL.key)
		assert.Equal(t, jan.high, rootLL.max)
		assert.Equal(t, "Jan", rootLL.payload)
		assert.True(t, tree.isLeaf(rootLL))

		assert.Equal(t, rootL, rootLR.parent)
		assert.Equal(t, black, rootLR.color)
		assert.Equal(t, mar, rootLR.key)
		assert.Equal(t, apr.high, rootLR.max)
		assert.Equal(t, "Mar", rootLR.payload)

		assert.Equal(t, rootLR, rootLRR.parent)
		assert.Equal(t, red, rootLRR.color)
		assert.Equal(t, apr, rootLRR.key)
		assert.Equal(t, apr.high, rootLRR.max)
		assert.Equal(t, "Apr", rootLRR.payload)
		assert.True(t, tree.isLeaf(rootLRR))

		assert.Equal(t, root, rootR.parent)
		assert.Equal(t, red, rootR.color)
		assert.Equal(t, nov, rootR.key)
		assert.Equal(t, dec.high, rootR.max)
		assert.Equal(t, "Nov", rootR.payload)

		assert.Equal(t, rootR, rootRL.parent)
		assert.Equal(t, black, rootRL.color)
		assert.Equal(t, jul, rootRL.key)
		assert.Equal(t, aug.high, rootRL.max)
		assert.Equal(t, "Jul", rootRL.payload)

		assert.Equal(t, rootRL, rootRLR.parent)
		assert.Equal(t, red, rootRLR.color)
		assert.Equal(t, aug, rootRLR.key)
		assert.Equal(t, aug.high, rootRLR.max)
		assert.Equal(t, "Aug", rootRLR.payload)
		assert.True(t, tree.isLeaf(rootRLR))

		assert.Equal(t, rootR, rootRR.parent)
		assert.Equal(t, black, rootRR.color)
		assert.Equal(t, dec, rootRR.key)
		assert.Equal(t, dec.high, rootRR.max)
		assert.Equal(t, "Dec", rootRR.payload)
		assert.True(t, tree.isLeaf(rootRR))

		tt := []struct {
			name       string
			query      Interval
			want       Result
			wantErrMsg string
		}{
			{
				name: "query unknown interval returns error",
				query: Interval{
					low:  newTime(t, "2019-Aug-01"),
					high: newTime(t, "2019-Aug-02"),
				},
				wantErrMsg: "no interval found",
			},
			{
				name: "query overlapping interval returns february",
				query: Interval{
					low:  newTime(t, "2020-Feb-01"),
					high: newTime(t, "2020-Feb-02"),
				},
				want: Result{
					Interval: feb,
					Payload:  "Feb",
				},
			},
			{
				name: "query big interval returns first hit",
				query: Interval{
					low:  newTime(t, "2020-Jan-01"),
					high: newTime(t, "2020-Dec-31"),
				},
				want: Result{
					Interval: may,
					Payload:  "May",
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				r, err := tree.FindFirstOverlapping(tc.query)

				if tc.wantErrMsg != "" {
					require.Error(t, err)
					assert.Contains(t, err.Error(), tc.wantErrMsg)
				} else {
					assert.NoError(t, err)
				}

				assert.Equal(t, tc.want, r)
			})
		}
	})
}
