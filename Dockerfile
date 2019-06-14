FROM redis:5.0.5-alpine

COPY redis.conf /etc/redis.conf

CMD ["redis-server","/etc/redis.conf"]