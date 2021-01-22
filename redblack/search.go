package redblack

// Result is a search result when looking up a Key in the tree.
type Result struct {
	Key     Key
	Payload interface{}
}

// InOrder returns an ordered list of all entries.
func (t *Tree) InOrder() []Result {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == t.sentinel {
		return nil
	}

	res := make([]Result, 0)

	t.resultsInorder(t.root, &res)

	return res
}

func (t *Tree) resultsInorder(z *node, res *[]Result) {
	if z == t.sentinel {
		return
	}

	if z.left != t.sentinel {
		t.resultsInorder(z.left, res)
	}

	*res = append(*res, Result{
		Key:     z.key,
		Payload: z.payload,
	})

	if z.right != t.sentinel {
		t.resultsInorder(z.right, res)
	}
}
