FROM ubuntu:16.04
COPY zips.csv /zips.csv
COPY zipsvr /zipsvr
ADD default.conf /etc/nginx/conf.d/default.conf
EXPOSE 443
ENTRYPOINT ["/zipsvr"]