build:
	@echo "Building application..."
	@go build -o bin/snapbin ./...
	@echo "Application build successful!\n"
run: build
	@./bin/snapbin