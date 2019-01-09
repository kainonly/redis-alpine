## Redis-Alpine

Redis minimalist custom image

- size `6.56` mb
- version `4.0.12`

Docker Pull Command

```shell
docker pull kainonly/redis-alpine
```

Set docker-compose

```yaml
version: '3'
services:
  redis:
    image: kainonly/redis-alpine:4.0.12
    restart: always
    privileged: true
    sysctls:
      net.core.somaxconn: 65535
    volumes:
      - ./redis/data:/data
    ports:
      - 6379:6379
```