GO := go
TEST_ARGS := -race -cover -short

test:
	$(GO) test $(TEST_ARGS) ./...