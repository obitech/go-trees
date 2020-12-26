package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTree_FindAllOverlapping(t *testing.T) {
	t.Run("search on empty tree returns error", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))

		r, err := tree.FindAllOverlapping(nov)
		assert.Error(t, err)
		assert.Nil(t, r)
	})

	t.Run("unknown interval search on rooted tree returns error", func(t *testing.T) {
		tree := NewIntervalTree()

		nov, _ := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
		dec, _ := NewInterval(newTime(t, "2020-Dec-01"), newTime(t, "2020-Dec-02"))

		tree.Upsert(dec, "Dec")

		r, err := tree.FindAllOverlapping(nov)
		assert.Error(t, err)
		assert.Nil(t, r)
	})

	t.Run("match on rooted tree returns result", func(t *testing.T) {
		tree := NewIntervalTree()

		dec, _ := NewInterval(newTime(t, "2020-Dec-01"), newTime(t, "2020-Dec-02"))

		tree.Upsert(dec, "Dec")

		r, err := tree.FindAllOverlapping(dec)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		want := []Result{
			{
				Interval: dec,
				Payload:  "Dec",
			},
		}

		assert.Equal(t, want, r)
	})

	t.Run("(h=1) match on root returns result", func(t *testing.T) {
		tree := NewIntervalTree()

		aug, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		sep, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))
		may, _ := NewInterval(newTime(t, "2020-May-01"), newTime(t, "2020-May-02"))

		tree.Upsert(aug, nil)
		tree.Upsert(sep, nil)
		tree.Upsert(may, nil)

		r, err := tree.FindAllOverlapping(aug)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		want := []Result{
			{
				Interval: aug,
			},
		}

		assert.Equal(t, want, r)
	})

	t.Run("(h=1) match on leaf returns result", func(t *testing.T) {
		tree := NewIntervalTree()

		aug, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		sep, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-02"))
		may, _ := NewInterval(newTime(t, "2020-May-01"), newTime(t, "2020-May-02"))

		tree.Upsert(aug, nil)
		tree.Upsert(sep, nil)
		tree.Upsert(may, nil)

		t.Run("find left leaf", func(t *testing.T) {
			r, err := tree.FindAllOverlapping(may)
			assert.NoError(t, err)
			assert.NotNil(t, r)

			want := []Result{
				{
					Interval: may,
				},
			}

			assert.Equal(t, want, r)
		})

		t.Run("find right leaf", func(t *testing.T) {
			r, err := tree.FindAllOverlapping(sep)
			assert.NoError(t, err)
			assert.NotNil(t, r)

			want := []Result{
				{
					Interval: sep,
				},
			}

			assert.Equal(t, want, r)
		})
	})

	t.Run("(h=1) searching for multiple overlaps returns them in order", func(t *testing.T) {
		tree := NewIntervalTree()

		ov1, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Oct-15"))
		ov2, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Sep-15"))
		ov3, _ := NewInterval(newTime(t, "2020-May-01"), newTime(t, "2020-Sep-02"))

		tree.Upsert(ov1, "ov1")
		tree.Upsert(ov2, "ov2")
		tree.Upsert(ov3, "ov3")

		s, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Sep-01"))

		r, err := tree.FindAllOverlapping(s)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		want := []Result{
			{
				Interval: ov3,
				Payload:  "ov3",
			},
			{
				Interval: ov1,
				Payload:  "ov1",
			},
			{
				Interval: ov2,
				Payload:  "ov2",
			},
		}

		assert.Equal(t, want, r)
	})

	t.Run("(h=3) searching for multiple overlaps returns them in order", func(t *testing.T) {
		tree := NewIntervalTree()

		search, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Oct-15"))

		nov, _ := NewInterval(newTime(t, "2020-Sep-01"), newTime(t, "2020-Nov-01")) // 3
		feb, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Dec-01")) // 2
		dec, _ := NewInterval(newTime(t, "2020-Dec-01"), newTime(t, "2020-Dec-02"))
		mar, _ := NewInterval(newTime(t, "2020-Mar-01"), newTime(t, "2020-Mar-02"))
		jan, _ := NewInterval(newTime(t, "2019-Jan-01"), newTime(t, "2021-Jan-02")) // 1
		jul, _ := NewInterval(newTime(t, "2020-Jul-01"), newTime(t, "2020-Jul-02"))
		may, _ := NewInterval(newTime(t, "2020-May-01"), newTime(t, "2020-May-02"))
		aug, _ := NewInterval(newTime(t, "2020-Aug-01"), newTime(t, "2020-Aug-02"))
		apr, _ := NewInterval(newTime(t, "2020-Apr-01"), newTime(t, "2020-Apr-02"))
		oct, _ := NewInterval(newTime(t, "2020-Oct-01"), newTime(t, "2020-Dec-15"))  // 4
		oct2, _ := NewInterval(newTime(t, "2020-Oct-14"), newTime(t, "2021-Dec-15")) // 5

		tree.Upsert(nov, "Nov")
		tree.Upsert(feb, "Feb")
		tree.Upsert(dec, "Dec")
		tree.Upsert(mar, "Mar")
		tree.Upsert(jan, "Jan")
		tree.Upsert(jul, "Jul")
		tree.Upsert(may, "May")
		tree.Upsert(aug, "Aug")
		tree.Upsert(apr, "Apr")
		tree.Upsert(oct, "Oct")
		tree.Upsert(oct2, "Oct21")

		r, err := tree.FindAllOverlapping(search)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		want := []Result{
			{
				Interval: jan,
				Payload:  "Jan",
			},
			{
				Interval: feb,
				Payload:  "Feb",
			},
			{
				Interval: nov,
				Payload:  "Nov",
			},
			{
				Interval: oct,
				Payload:  "Oct",
			},
			{
				Interval: oct2,
				Payload:  "Oct21",
			},
		}

		assert.Equal(t, want, r)
	})
}
