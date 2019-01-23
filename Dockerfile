FROM redis:alpine

COPY redis.conf /etc/redis.conf

VOLUME /data

EXPOSE 6379

CMD ["redis-server","/etc/redis.conf"]