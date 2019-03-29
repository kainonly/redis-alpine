## Redis-Alpine

Docker Pull Command 

```shell
docker pull kainonly/redis-alpine
```

Set docker-compose

```yaml
version: '3.7'
services:
  redis:
    image: kainonly/redis-alpine
    restart: always
    privileged: true
    sysctls:
      net.core.somaxconn: 65535
    volumes:
      - ./redis/redis.conf:/etc/redis.conf
      - /home/redis/data:/data
    ports:
      - 6379:6379
```
