FROM redis:6.0.9-alpine

COPY redis.conf /etc/redis.conf

CMD ["redis-server","/etc/redis.conf"]