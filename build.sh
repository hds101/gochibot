#!/bin/bash

CURRENT_PATH="$( cd "$( dirname "$(readlink -f $0)" )" && pwd )"
cd $CURRENT_PATH/src

GOOS=linux GOARCH=amd64 go get
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $CURRENT_PATH/bin/gochibot
