# Makefile

.PHONY: build clean fmt vet

# Set the binary name
BINARY=plb

# Format and vet the Go application
fmt:
	go fmt ./...

vet:
	go vet ./...

# Build the Go application
build:
	go build -o out/$(BINARY) .

# Clean up the binary
clean:
	rm -f out/$(BINARY)
