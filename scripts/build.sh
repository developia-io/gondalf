#!/usr/bin/env bash

bin_name="gondalf"

mkdir -p bin/linux/amd64
mkdir -p bin/linux/arm64
mkdir -p bin/darwin/amd64
mkdir -p bin/darwin/arm64
mkdir -p bin/windows/amd64
mkdir -p bin/windows/386

GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o "bin/linux/amd64/$bin_name"
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o "bin/linux/arm64/$bin_name"
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o "bin/darwin/amd64/$bin_name"
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o "bin/darwin/arm64/$bin_name"
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o "bin/windows/amd64/$bin_name.exe"
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -trimpath -o "bin/windows/386/$bin_name.exe"

tar czvf bin/linux/amd64/$bin_name-linux-amd64.tar.gz -C bin/linux/amd64 $bin_name
tar czvf bin/linux/arm64/$bin_name-linux-arm64.tar.gz -C bin/linux/arm64 $bin_name
tar czvf bin/darwin/amd64/$bin_name-darwin-amd64.tar.gz -C bin/darwin/amd64 $bin_name
tar czvf bin/darwin/arm64/$bin_name-darwin-arm64.tar.gz -C bin/darwin/arm64 $bin_name
tar czvf bin/windows/amd64/$bin_name-windows-amd64.tar.gz -C bin/windows/amd64 $bin_name.exe
tar czvf bin/windows/386/$bin_name-windows-386.tar.gz -C bin/windows/386 $bin_name.exe
