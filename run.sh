#!/bin/bash

MODULE_NAME="Zuxyll-game"

if [ -f "go.mod" ]; then
    echo "go.mod already exists. Checking for updates..."
    go get -u
else
    echo "go.mod not found. Initializing Go module..."
    go mod init "$MODULE_NAME"
    echo "Go module initialized with name: $MODULE_NAME"
fi

go mod tidy
go run .
