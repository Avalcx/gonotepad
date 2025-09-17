#!/bin/bash
cd "$(pwd)" || exit

export NAME="notepad"

#x64
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o $NAME-$GOOS-$GOARCH
