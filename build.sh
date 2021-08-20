#!/bin/bash

# Build binaries for Windows 64bit and various Linux architectures

archs=(amd64 arm arm64 386)

echo "Building Win64 executable"
go build -o build/ip-checker-win64.exe .

for arch in "${archs[@]}"; do
  echo "Building Linux/$arch binary"
  GOOS=linux GOARCH=$arch go build -o build/ip-checker-"$arch" .
done

# Build zip for the Lambda function

echo "Building Lambda function zip"
cp .env ip.txt build
zip -rjq build/ip-checker.zip build/ip-checker-amd64 build/.env build/ip.txt
rm build/.env build/ip.txt
