
# Go Port Proxy Server

This Go application sets up a simple HTTP server that listens on port `2701` and proxies requests to various ports on `127.0.0.1`. It allows you to access services running on different ports by using a URL path pattern.

## Features

- Forwards HTTP requests from `http://127.0.0.1:2701/{port}` to `http://127.0.0.1:{port}`.
- Handles all ports from `0` to `65535`.
- Useful for accessing local services through a single entry point.

## Requirements

- Go (version 1.16 or higher)

## Installation

# Auto Installation

The `install.sh` script simplifies the installation process by handling environment detection, Go installation, and building the `portproxy` application. It supports both Linux and Termux environments.

1. **Clone the Repository**

   ```bash
   git clone https://github.com/mahendraplus/Go-Port-Proxy-Server
   cd Go-Port-Proxy-Server
   ```

2. **Run the Installation Script**

   Make the `install.sh` script executable and run it:

   ```bash
   chmod +x install.sh
   ./install.sh
   ```

   - **For Termux:** Installs Go and places the `portproxy` binary in `/data/data/com.termux/files/usr/bin/`.
   - **For Linux:** Installs Go and places the `portproxy` binary in `/usr/local/bin/`.

3. **Run the Application**

   After installation, you can start the proxy server by simply typing:

   ```bash
   portproxy
   ```


## Manually 

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mahendraplus/Go-Port-Proxy-Server
   cd Go-Port-Proxy-Server
   ```

2. **Build the Go application:**

   ```bash
   go build -o portproxy.go
   ```


3. **Run the proxy server:**

   ```bash
   ./portproxy
   ```

   This will start the server on port `2701`.

 **Access a service running on port `8080` on your local machine by visiting:**

   ```text
   http://127.0.0.1:2701/8080
   ```

   This URL will proxy the request to `http://127.0.0.1:8080`.

## Example: Accessing SSH Through Proxy

To forward SSH traffic through the proxy server, follow these steps:

1. **Run the Go proxy server on port `2701`.**

2. **Set up local port forwarding via SSH to use the proxy:**

   If you want to connect to an SSH service on port `22` on `127.0.0.1` but access it through a proxy listening on port `2701`, you can do this by setting up a local port forwarding:

   ```bash
   ssh -L 2701:127.0.0.1:22 -p 8022 user@remote_host
   ```

   - `-L 2701:127.0.0.1:22`: Forwards local port `2701` to port `22` on `127.0.0.1`.
   - `-p 8022`: Specifies the port on the SSH server to connect to.
   - `user@remote_host`: Replace `user` and `remote_host` with your SSH username and the remote server address.

3. **Connect to the SSH service using:**

   ```bash
   ssh -p 2701 your_ssh_user@127.0.0.1
   ```

   This command connects to the SSH service on the remote server via the local port `2701`.

## YE BHI READ KARLO....

- Ensure that the port you are forwarding to is not already in use by another service.
- Make sure that your firewall and security settings allow traffic on the necessary ports.
