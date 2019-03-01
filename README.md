## Redis-Alpine

Redis custom image

Docker Pull Command

```shell
docker pull kainonly/redis-alpine:5.0.3
```

Set docker-compose

```yaml
version: '3'
services:
  redis:
    image: kainonly/redis-alpine:5.0.3
    restart: always
    privileged: true
    sysctls:
      net.core.somaxconn: 65535
    volumes:
      - ./redis/data:/data
    ports:
      - 6379:6379
```
