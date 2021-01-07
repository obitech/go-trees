package interval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func newTime(t *testing.T, s string) time.Time {
	x, err := time.Parse("2006-Jan-02", s)
	require.NoError(t, err)

	return x
}

func TestInterval_less(t *testing.T) {
	tt := []struct {
		name string
		x    Interval
		y    Interval
		want bool
	}{
		// x |
		// y |
		{
			name: "less against empty Interval returns false",
			x:    Interval{},
		},
		// x   |---|
		// y |---|
		{
			name: "x greater y: x high and low greater y returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-01"),
				high: newTime(t, "2020-Jan-04"),
			},
		},

		// x       |---|
		// y |---|
		{
			name: "x greater y: x high and low greater y returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-05"),
				high: newTime(t, "2020-Jan-06"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-01"),
				high: newTime(t, "2020-Jan-04"),
			},
		},
		{
			name: "Comparing two default intervals returns false",
			x:    Interval{},
			y:    Interval{},
		},
		// x |---|
		// y        |---|
		{
			name: "x smaller y: no overlap",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-04"),
			},
			y: Interval{
				low:  newTime(t, "2020-Feb-02"),
				high: newTime(t, "2020-Feb-03"),
			},
			want: true,
		},
		// x |---|
		// y   |---|
		{
			name: "x smaller y: overlap",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-04"),
				high: newTime(t, "2020-Jan-06"),
			},
			want: true,
		},
		// x |---|
		// y |-----|
		{
			name: "x smaller y: same low, smaller high",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-06"),
			},
			want: true,
		},
		// x |---|
		// y |---|
		{
			name: "x equal y returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
		},
		// x |-----|
		// y |---|
		{
			name: "x greater y: x high greater y high returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-04"),
			},
		},
		// x   |---|
		// y |---|
		{
			name: "x greater y: x high and low greater y returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-01"),
				high: newTime(t, "2020-Jan-04"),
			},
		},

		// x       |---|
		// y |---|
		{
			name: "x greater y: x high and low greater y returns false",
			x: Interval{
				low:  newTime(t, "2020-Jan-05"),
				high: newTime(t, "2020-Jan-06"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-01"),
				high: newTime(t, "2020-Jan-04"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.x.less(tc.y))
		})
	}
}

func TestInterval_overlaps(t *testing.T) {
	tt := []struct {
		name string
		x    Interval
		y    Interval
		want bool
	}{
		// x |---|
		// y       |---|
		{
			name: "x < y returns false",
			x: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Mar-02"),
			},
			y: Interval{
				low:  newTime(t, "2020-Apr-01"),
				high: newTime(t, "2020-Apr-02"),
			},
		},
		// x       |---|
		// y |---|
		{
			name: "x > y returns false",
			x: Interval{
				low:  newTime(t, "2020-Apr-01"),
				high: newTime(t, "2020-Apr-02"),
			},
			y: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Mar-02"),
			},
		},
		// x  |--|
		// y |----|
		{
			name: "x inside y returns true",
			x: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Mar-02"),
			},
			y: Interval{
				low:  newTime(t, "2020-Feb-01"),
				high: newTime(t, "2020-Apr-02"),
			},
			want: true,
		},
		// x |----|
		// y  |--|
		{
			name: "y inside x returns true",
			x: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Aug-01"),
			},
			y: Interval{
				low:  newTime(t, "2020-Apr-01"),
				high: newTime(t, "2020-May-01"),
			},
			want: true,
		},
		// x |---|
		// y  |----|
		{
			name: "x.high inside inside y returns true",
			x: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Apr-01"),
			},
			y: Interval{
				low:  newTime(t, "2020-Mar-25"),
				high: newTime(t, "2020-May-01"),
			},
			want: true,
		},
		// x    |---|
		// y |----|
		{
			name: "x.low inside inside y returns true",
			x: Interval{
				low:  newTime(t, "2020-Mar-01"),
				high: newTime(t, "2020-Apr-01"),
			},
			y: Interval{
				low:  newTime(t, "2020-Feb-01"),
				high: newTime(t, "2020-Mar-25"),
			},
			want: true,
		},
		// x     |---|
		// y |---|
		{
			name: "y.high equal x.low returns true",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-01"),
				high: newTime(t, "2020-Jan-03"),
			},
			want: true,
		},
		// x |---|
		// y     |---|
		{
			name: "x.high equal y.low returns true",
			x: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			y: Interval{
				low:  newTime(t, "2020-Jan-05"),
				high: newTime(t, "2020-Jan-06"),
			},
			want: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.x.overlaps(tc.y))
		})
	}
}

func TestInterval_intersects(t *testing.T) {
	tt := []struct {
		name string
		i    Interval
		t    time.Time
		want bool
	}{
		{
			name: "t before i returns false",
			i: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			t: newTime(t, "2019-Jan-01"),
		},
		{
			name: "t after i returns false",
			i: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			t: newTime(t, "2021-Jan-01"),
		},
		{
			name: "t == i.low returns true",
			i: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			t:    newTime(t, "2020-Jan-03"),
			want: true,
		},
		{
			name: "t == i.high returns true",
			i: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			t:    newTime(t, "2020-Jan-05"),
			want: true,
		},
		{
			name: "t inside i returns true",
			i: Interval{
				low:  newTime(t, "2020-Jan-03"),
				high: newTime(t, "2020-Jan-05"),
			},
			t:    newTime(t, "2020-Jan-04"),
			want: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.i.intersects(tc.t))
		})
	}
}

func Test_greaterOrEqual(t *testing.T) {
	tt := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "equal times returns true",
			t1:   newTime(t, "2020-Jan-01"),
			t2:   newTime(t, "2020-Jan-01"),
			want: true,
		},
		{
			name: "t1 > t2 returns true",
			t1:   newTime(t, "2020-Jan-02"),
			t2:   newTime(t, "2020-Jan-01"),
			want: true,
		},
		{
			name: "t1 < t2 returns false",
			t1:   newTime(t, "2020-Jan-01"),
			t2:   newTime(t, "2020-Jan-02"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, greaterOrEqual(tc.t1, tc.t2))
		})
	}
}
