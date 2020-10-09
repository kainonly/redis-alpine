FROM redis:6.0.8-alpine

COPY redis.conf /etc/redis.conf

CMD ["redis-server","/etc/redis.conf"]