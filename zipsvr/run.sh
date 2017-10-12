#!/usr/bin/env bash
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
# -v /etc/letsencrypt:/etc/letsencrypt:ro \
# -e TLSCERT=$TLSCERT \
# -e TLSKEY=$TLSKEY \
# tzchen/zipsvr
-v $(pwd)/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
tzchen/zipsvr