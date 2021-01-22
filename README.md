# go-tree: Tree implementations in Golang

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-82%25-brightgreen.svg?longCache=true&style=flat)</a>

## package [bst](./bst)

Implements a [Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree).

## package [redblack](./redblack)

Implements a [Red-Black-Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree).

### Example

```go
tree := NewRedBlackTree()

// Operations on arbitrary keys are allow as long as the type implements they 
// Key interface.
type myInt int

func (i myInt) Less(v Key) bool {
    return i < v.(myInt)
}

// Insert items.
tree.Upsert(myInt(5), "test")
tree.Upsert(myInt(10), "foo")
tree.Upsert(myInt(15), 42)

// Search for an item.
fmt.Println(tree.Search(15)) 

// Replace a payload.
tree.Upsert(myInt(15), "bar")
```

### Benchmarks

````
goos: darwin
goarch: amd64
pkg: github.com/obitech/go-tree/redblack
BenchmarkRBTree_Upsert-8                 2012917               775 ns/op              31 B/op          0 allocs/op
BenchmarkRBTree_Search10_000-8          11450182               102 ns/op               0 B/op          0 allocs/op
BenchmarkRBTree_Search100_000-8          5394513               235 ns/op               0 B/op          0 allocs/op
BenchmarkRBTree_Search1_000_000-8         411782              2487 ns/op             117 B/op          3 allocs/op
BenchmarkRBTree_Delete10_000-8          25755091                43.7 ns/op             0 B/op          0 allocs/op
BenchmarkRBTree_Delete100_000-8         23051653                46.5 ns/op             0 B/op          0 allocs/op
BenchmarkRBTree_Delete1_000_000-8          37518             26899 ns/op            1291 B/op         43 allocs/op
````

## package [interval](./interval)

Implements an [Interval tree](https://en.wikipedia.org/wiki/Interval_tree)

### Example

```go
tree := NewIntervalTree()

nov, err := NewInterval(newTime(t, "2020-Nov-01"), newTime(t, "2020-Nov-02"))
if err != nil {
	panic(err)
}


feb, err := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2020-Feb-02"))
if err != nil {
	panic(err)
}

dec, err := NewInterval(newTime(t, "2020-Dec-01"), newTime(t, "2020-Dec-02"))
if err != nil {
	panic(err)
}

tree.Upsert(nov)
tree.Upsert(feb)
tree.Upsert(dec)

search1, err := NewInterval(newTime(t, "2020-Feb-01"), newTime(t, "2021-Feb-02"))

r, err := tree.FindFirstOverlapping(search)
if err != nil {
	fmt.Println("no overlapping interval found for %s", search1)
}

r, err = tree.FindAllOverlapping(search)
if err != nil {
    fmt.Println("no overlapping interval found for %s", search1)
}
```

### Benchmarks

TODO