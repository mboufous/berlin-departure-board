server:
	@go run ./cmd/server/main.go
test:
	@go test -v ./...

debug-test-ai:
	@copilot-debug go test -v ./...

.PHONY: server test