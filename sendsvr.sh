#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o apiserver .
scp apiserver .env ziumks@zium01:~/apiserver
