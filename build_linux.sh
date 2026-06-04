#!/bin/bash
echo "Building Labelin for Linux..."
export GOOS=linux
export GOARCH=amd64
go build -o labelin_linux_amd64 cmd/server/main.go
echo "Build complete."
