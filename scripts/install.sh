#!/usr/bin/env bash

version="0.0.1"
os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

if [ "$os" != "linux" ] && [ "$os" != "darwin" ]; then
    echo "Unsupported OS: $os"
    exit 1
elif [ "$arch" != "amd64" ] && [ "$arch" != "arm64" ] && [ "$arch" != "aarch64" ] && [ "$arch" != "x86_64" ]; then
    echo "Unsupported architecture: $arch"
    exit 1
fi

if [ "$arch" == "aarch64" ]; then
    arch="arm64"
elif [ "$arch" == "x86_64" ]; then
    arch="amd64"
fi
  
echo "Downloading and installing Gondalf $version for $os/$arch..."

url="https://github.com/developia-io/gondalf/releases/download/v$version/gondalf-$os-$arch.tar.gz"
curl -sL "$url" | sudo tar xz -C /usr/local/bin

echo "Gondalf $version has been installed successfully!"
