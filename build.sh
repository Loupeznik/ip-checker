#!/bin/bash

# Build binaries for Windows 64bit and various Linux architectures

archs=(amd64 arm arm64 386)

echo "Building Win64 executable"
go build -o build/ip-checker-win64.exe .

for arch in "${archs[@]}"; do
  echo "Building Linux/$arch binary"
  GOOS=linux GOARCH=$arch go build -o build/ip-checker-"$arch" .
done

# Copy environment variables and the current IP file to accompany the built binaries

echo "Moving files"
cp .env ip.txt build
