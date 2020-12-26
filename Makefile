GO := go
TEST_ARGS := -race -cover

test:
	$(GO) test $(TEST_ARGS) ./...

bench: bench-rbt

bench-rbt:
	cd redblack/ && $(GO) test -bench=. -benchmem

report:
	$(GO) test -cover -coverprofile=cover.out ./...
	$(GO) tool cover -html=cover.out

coverage-badge:
	gopherbadger -md="README.md"