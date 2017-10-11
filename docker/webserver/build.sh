#!/usr/bin/env bash
set -e
go build
docker build -t tzchen/testserver .
docker push tzchen/testserver