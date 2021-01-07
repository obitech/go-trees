package interval

import "time"

type node struct {
	key     Interval
	color   color
	left    *node
	right   *node
	parent  *node
	max     time.Time
	payload interface{}
}
