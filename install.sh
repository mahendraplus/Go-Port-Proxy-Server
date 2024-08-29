#!/bin/bash


# ©mahendraplus [29/08/2024] (MAX)
# This script installs Go, builds the Go application from portproxy.go, and moves it to a PATH directory.

# Function to install Go on Linux
install_go_linux() {
    local GO_VERSION="1.20.5"  # Update this to the latest version as needed
    local GO_URL="https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz"
    local GO_DIR="/usr/local/go"

    echo "Installing Go ${GO_VERSION} for Linux..."
    wget "${GO_URL}" -O go.tar.gz
    sudo tar -C /usr/local -xzf go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo "Go installed. Version: $(go version)"
}

# Function to install Go on Termux
install_go_termux() {
    echo "Installing Go for Termux..."
    pkg install -y golang
    echo "Go installed. Version: $(go version)"
}

# Detect the environment and install Go accordingly
if [ "$TERMUX_VERSION" ]; then
    echo "Running in Termux environment."
    install_go_termux
elif [ -f /etc/os-release ]; then
    echo "Running in Linux environment."
    install_go_linux
else
    echo "Unsupported environment. Please install Go manually."
    exit 1
fi

# Clone the repository if not already cloned
if [ ! -d "portproxy" ]; then
    echo "Cloning the Go proxy repository..."
    git clone https://github.com/mahendraplus/Go-Port-Proxy-Server
fi

# Navigate to the repository directory
cd Go-Port-Proxy-Server || exit

# Build the Go application
echo "Building the Go proxy application..."
go build -o portproxy portproxy.go

# Move the binary to a directory in PATH
if [ "$TERMUX_VERSION" ]; then
    echo "Moving binary to /data/data/com.termux/files/usr/bin/"
    mv portproxy /data/data/com.termux/files/usr/bin/portproxy
elif [ -f /etc/os-release ]; then
    echo "Moving binary to /usr/local/bin/"
    sudo mv portproxy /usr/local/bin/portproxy
else
    echo "Unsupported environment. Please move the binary manually."
    exit 1
fi

echo "Installation completed successfully. You can now run 'portproxy' from anywhere."
