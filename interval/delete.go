package interval

// Delete deletes a node with the given key.
func (t *Tree) Delete(key Interval) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.root == t.sentinel {
		return
	}

	if n := t.findExact(key); n != nil {
		t.delete(n)
	}
}

func (t *Tree) delete(z *node) {
	var (
		// y is either removed or moved in tree
		y = z
		// If the color changes, we need to fix it
		yOriginalColor = y.color
		// This node moves into y's original position
		x *node
	)

	switch {
	case z.left == t.sentinel:
		x = z.right
		t.transplant(z, z.right)
	case z.right == t.sentinel:
		x = z.left
		t.transplant(z, z.left)
	default:
		// y is now the successor of z
		y = t.min(z.right)
		yOriginalColor = y.color

		// x will move into z's position
		x = y.right

		// If the successor is the direct children of z, update relationship
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)

		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	t.recalcMax(x)

	if yOriginalColor == black {
		t.fixupDelete(x)
	}
}

func (t *Tree) transplant(u, v *node) {
	switch {
	case u.parent == t.sentinel:
		t.root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}

	v.parent = u.parent
}

func (t *Tree) fixupDelete(x *node) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right

			if w.color == red {
				w.color = black
				x.parent.color = red

				t.rotateLeft(x.parent)

				w = x.parent.right
			}

			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red

					t.rotateRight(w)

					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black

				t.rotateLeft(x.parent)

				x = t.root
			}
		} else {
			w := x.parent.left

			if w.color == red {
				w.color = black
				x.parent.color = red

				t.rotateRight(x.parent)

				w = x.parent.left
			}

			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red

					t.rotateLeft(w)

					w = x.parent.left
				}

				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black

				t.rotateRight(x.parent)

				x = t.root
			}
		}
	}
	x.color = black
}
