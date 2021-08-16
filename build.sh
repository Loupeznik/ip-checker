#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o build/ip-checker .
cp .env ip.txt build
zip -rj build/ip-checker.zip build/ip-checker build/.env build/ip.txt
rm build/.env build/ip.txt
