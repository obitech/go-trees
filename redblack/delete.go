package redblack

// Delete deletes a node with the given key.
func (t *Tree) Delete(key Key) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if n := t.search(t.root, key); n != t.sentinel {
		t.delete(n)
	}
}

func (t *Tree) delete(z *node) {
	var (
		y              = z
		yOriginalColor = y.color
		x              *node
	)

	switch {
	case z.left == t.sentinel:
		x = z.right
		t.transplant(z, z.right)
	case z.right == t.sentinel:
		x = z.left
		t.transplant(z, z.left)
	default:
		y = t.min(z.right)
		yOriginalColor = y.color

		x = y.right

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

			switch {
			case w.left.color == black && w.right.color == black:
				w.color = red
				x = x.parent
			case w.right.color == black:
				w.left.color = black
				w.color = red

				t.rotateRight(w)

				w = x.parent.right
			default:
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

			switch {
			case w.right.color == black && w.left.color == black:
				w.color = red
				x = x.parent
			case w.left.color == black:
				w.right.color = black
				w.color = red

				t.rotateLeft(w)

				w = x.parent.left
			default:
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
