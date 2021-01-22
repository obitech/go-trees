package redblack

// Upsert updates an existing payload, or inserts a new one with the given key.
func (t *Tree) Upsert(key Key, payload interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if existing := t.search(t.root, key); existing != t.sentinel {
		existing.payload = payload
	} else {
		t.insert(t.newLeaf(key, payload))
	}
}

func (t *Tree) insert(z *node) {
	var (
		y = t.sentinel
		x = t.root
	)

	for x != t.sentinel {
		y = x
		if z.key.Less(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y

	switch {
	case y == t.sentinel:
		t.root = z
	case z.key.Less(y.key):
		y.left = z
	default:
		y.right = z
	}

	z.left = t.sentinel
	z.right = t.sentinel
	z.color = red

	t.fixupInsert(z)
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
