#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t tzchen/testserver .
docker push tzchen/testserver