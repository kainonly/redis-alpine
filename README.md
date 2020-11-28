## Redis Alpine

Configurable Redis Image

![MicroBadger Size](https://img.shields.io/microbadger/image-size/kainonly/redis-alpine.svg?style=flat-square)
![MicroBadger Layers](https://img.shields.io/microbadger/layers/kainonly/redis-alpine.svg?style=flat-square)
![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/redis-alpine.svg?style=flat-square)
[![Github Actions](https://img.shields.io/github/workflow/status/docker-maker/redis-alpine/release?style=flat-square)](https://github.com/docker-marker/redis-alpine/actions)

```shell
docker pull kainonly/redis-alpine
```

Set docker-compose

```yaml
version: '3.8'
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
