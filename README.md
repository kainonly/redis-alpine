## Redis Alpine

Configurable Redis Image

![MicroBadger Size](https://img.shields.io/microbadger/image-size/kainonly/redis-alpine.svg?style=flat-square)
![MicroBadger Layers](https://img.shields.io/microbadger/layers/kainonly/redis-alpine.svg?style=flat-square)
![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/redis-alpine.svg?style=flat-square)
![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/kainonly/redis-alpine.svg?style=flat-square)
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/kainonly/redis-alpine.svg?style=flat-square)

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
