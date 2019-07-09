#!/bin/bash
 
export GO111MODULE=on
go mod tidy
go mod verify
 
if [[ "$1" == "lint" ]]; then
    echo "#######################"
    echo "# Running GolangCI-Lint"
    golangci-lint --version
    golangci-lint run --enable-all --disable gochecknoglobals --disable gochecknoinits
fi
 
if [[ "$1" == "test" ]]; then
    echo "#######################"
    echo "# Running Test"
    go test ./... -race
fi
