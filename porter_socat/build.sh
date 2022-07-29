#!/bin/bash

root=$(dirname "$(greadlink -f "$0")")
mkdir -p "$root/build/macos"
cd "$root/build" || exit 1
env GOARCH=amd64 GOOS=linux go build ..
cd "$root/build/macos" || exit 1
go build ../../