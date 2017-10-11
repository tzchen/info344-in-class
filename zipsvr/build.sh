#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t tzchen/zipsvr .
docker push tzchen/zipsvr