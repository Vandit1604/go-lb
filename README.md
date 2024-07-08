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
   git clone https://github.com/Vandit1604/go-lb.git
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

1. Start the python testing-servers with following commands in indiviual terminals  

```
python3 -m http.server 8080 --directory testing-servers/server8080
python3 -m http.server 8081 --directory testing-servers/server8081
python3 -m http.server 8082 --directory testing-servers/server8082

```

2. Visit *http://localhost:9000* 

3. Do hard refresh with <kbd>CTRL+SHIFT+R</kbd> and you'll be able to see the port changing in HTML.

---
NOTE: Use this command to kill all the python servers after you're done testing
```
pkill -f "python3 -m http.server"
```
---

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
