package interval

import (
	"errors"
	"fmt"
	"time"
)

// Interval marks a timespan with a start and end time.
type Interval struct {
	low  time.Time
	high time.Time
}

// NewInterval returns a new Interval or an error if end is before start.
func NewInterval(start, end time.Time) (Interval, error) {
	if end.Before(start) {
		return Interval{}, errors.New("start must be before end")
	}

	return Interval{
		low:  start,
		high: end,
	}, nil
}

func (i Interval) less(x Interval) bool {
	return i.low.Before(x.low) || i.low == x.low && i.high.Before(x.high)
}

func (i Interval) overlaps(x Interval) bool {
	return (i.low.Equal(x.high) || i.low.Before(x.high)) && (x.low.Equal(i.high) || x.low.Before(i.high))
}

func (i Interval) intersects(t time.Time) bool {
	return (i.low.Equal(t) || i.low.Before(t)) && (i.high.Equal(t) || i.high.After(t))
}

func (i Interval) String() string {
	return fmt.Sprintf("{start: %s, end: %s}", i.low, i.high)
}

func greaterOrEqual(t1, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}
