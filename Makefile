build:
	@echo "Building application..."
	@go build -o bin/snapbin ./cmd/web
	@echo "Application build successful!\n"
run: build
	@./bin/snapbin -addr=$(addr)
help:
	@go run ./... -help