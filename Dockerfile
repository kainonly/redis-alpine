FROM alpine:edge

ENV VERSION 4.0.12

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache redis

COPY redis.conf /etc/redis.conf

VOLUME /data

EXPOSE 6379

STOPSIGNAL SIGTERM

CMD ["redis-server","/etc/redis.conf"]