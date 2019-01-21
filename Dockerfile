FROM redis:alpine

ENV VERSION 5.0.3

COPY redis.conf /etc/redis.conf

VOLUME /data

EXPOSE 6379

STOPSIGNAL SIGTERM

CMD ["redis-server","/etc/redis.conf"]