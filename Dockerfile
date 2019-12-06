FROM redis:5.0.7-alpine

COPY redis.conf /etc/redis.conf

CMD ["redis-server","/etc/redis.conf"]