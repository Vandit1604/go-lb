# Simple Round-Robin Load Balancer

This project implements a simple round-robin load balancer in Go. The load balancer uses a configuration file (`config.yml`) to determine which backend servers to distribute incoming HTTP requests among.

## Features

- **Round-Robin Strategy**: Distributes incoming requests evenly across all available backend servers.
- **Configuration-Based**: Backend servers are configured through a `config.yml` file.

## Getting Started

### Prerequisites

Ensure you have Go installed on your machine. You can download it from [the official Go website](https://go.dev).

### Installation

1. Clone the repository to your local machine.
   ```bash
   git clone https://your-repository-url.git
   cd go-lb 
   ```

2. Install dependencies.
   ```bash
   go mod tidy
   ```

3. Configure the load balancer by editing `config.yml` to include your backend servers' URLs.

### Building the Load Balancer

Run the following command to build the load balancer executable.
```bash
make build
```
This command compiles the Go code into a binary named `go-lb` located in the `bin` directory.

### Starting the Load Balancer

Start the load balancer by executing the built binary.
```bash
./bin/go-lb
```

The load balancer listens on port 9000 by default. You can change the port by modifying the `port` field in the `config.yml` file.

## Testing the Load Balancer

To test the load balancer, you can use the provided Makefile targets.

1. **Start Multiple Test Servers**: Use the `make servers` command to start multiple test servers on different ports. This simulates a scenario where the load balancer receives requests from multiple clients.

2. **Build and Run the Load Balancer**: Execute `make run` to build and run the load balancer. It will start listening for incoming requests and distribute them among the test servers.

3. **Stop the Load Balancer and Test Servers**: Use `make stop` to gracefully shut down the load balancer and all test servers.

## Configuration

The load balancer's behavior is determined by the `config.yml` file. This file specifies the backend servers the load balancer should distribute requests to. Here's an example configuration:

```yaml
port: 3332
strategy: round-robin 
backends:
  - "http://localhost:8080"
  - "http://localhost:8081"
  - "http://localhost:8082"
```

Each entry in the `backends` list represents a backend server the load balancer can forward requests to.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and enhancements.
