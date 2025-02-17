.PHONY: test build run clean fmt vet bench

# Run all tests in the project
test:
	go test -v ./...

# Build the Go project
build:
	go build -o myapp ./...

# Run the Go application
run: build
	./myapp

# Format the code using go fmt
fmt:
	go fmt ./...

# Run go vet to catch potential issues
vet:
	go vet ./...

# Run Benchmark test
bench:
	go test -bench=. -benchmem

# Clean up generated binaries
clean:
	rm -f myapp
