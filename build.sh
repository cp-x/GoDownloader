#!/bin/bash

mkdir -p bin
archs=("386" "amd64" "arm64" "mips" "ppc64")

for arch in ${archs[@]}; do
    GOOS=linux GOARCH=$arch go build -o bin/go-downloader-$arch
done