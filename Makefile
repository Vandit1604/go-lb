# Build the Go load balancer binary
build:
	@go build -o bin/go-lb

# Run the Go load balancer after building
run: build
	@./bin/go-lb

