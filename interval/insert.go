package interval

// Upsert updates an existing payload, or inserts a new one with the given
// interval key.
func (t *Tree) Upsert(key Interval, payload interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.root == t.sentinel {
		t.insert(t.newLeaf(key, payload))

		return
	}

	if n := t.findExact(key); n != nil {
		n.payload = payload
	} else {
		t.insert(t.newLeaf(key, payload))
	}
}

func (t *Tree) insert(z *node) {
	var (
		y = t.sentinel
		x = t.root
	)

	// Find node to attach it to.
	for x != t.sentinel {
		y = x
		if z.key.less(x.key) {
			x = x.left
		} else {
			x = x.right
		}

		// Update max as we walk.
		if z.max.After(y.max) {
			y.max = z.max
		}
	}

	z.parent = y

	switch {
	case y == t.sentinel:
		t.root = z
	case z.key.less(y.key):
		y.left = z
	default:
		y.right = z
	}

	z.left = t.sentinel
	z.right = t.sentinel
	z.color = red

	t.fixupInsert(z)
}

func (t *Tree) recalcMax(z *node) {
	for z != t.sentinel {
		t.updateMax(z)
		z = z.parent
	}
}

func (t *Tree) fixupInsert(z *node) {
	for z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right

			switch {
			case y.color == red:
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			case z == z.parent.right:
				z = z.parent
				t.rotateLeft(z)
			default:
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left

			switch {
			case y.color == red:
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			case z == z.parent.left:
				z = z.parent
				t.rotateRight(z)
			default:
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateLeft(z.parent.parent)
			}
		}
	}
	t.root.color = black
}
