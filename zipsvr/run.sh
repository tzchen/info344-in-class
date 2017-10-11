#!/usr/bin/env bash
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v ~/UW/INFO344/go/src/github.com/tzchen/info344-in-class/zipsvr/tls:tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
tzchen/zipsvr