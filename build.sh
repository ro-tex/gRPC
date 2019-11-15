#!/usr/bin/env bash
#CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -o server/server ./server
#CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -o client/client ./client
CGO_ENABLED=0 GOOS=linux go build -o server/server ./server
CGO_ENABLED=0 GOOS=linux go build -o client/client ./client

