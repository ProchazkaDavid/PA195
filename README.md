# PA195 NoSQL Databases - Redis

> [NoSQL Databases, Fall 2019 – Group Projects](http://disa.fi.muni.cz/vlastislav-dohnal/teaching/nosql-databases-fall-2019/group-projects/)

## Running a single instance

```bash
# Create a network for a Redis
docker network create redisnet

# Run a single instance that is automatically removed after being killed
docker run --rm --net redisnet --name redis -d redis

# Now you should see Redis container running
docker ps

# Launch redis-cli that is automatically removed after exit
docker run --rm --net redisnet --name redis-cli -it goodsmileduck/redis-cli redis-cli -h redis.redisnet

# After you are done with experimenting, kill Redis container
docker kill redis
```

> NOTE: You can add `-p 127.0.0.1:6379:6379` to your redis container to access port locally and not only through redis-cli in the docker container. 

## Running a cluster

[Redis cluster tutorial](https://redis.io/topics/cluster-tutorial)

### Starting a cluster 

```bash
docker-compose up
```

### Initialization

```bash
docker exec -it redis-1 redis-cli -p 7000 --cluster create \
192.168.0.2:7000 \
192.168.0.3:7001 \
192.168.0.4:7002 \
192.168.0.5:7003 \
192.168.0.6:7004 \
192.168.0.7:7005 \
--cluster-replicas 1
```

Then type `yes` and you are ready to go.

### Connecting to the cluster

> NOTE: use `-c` to enable cluster mode

```bash
docker exec -it redis-1 redis-cli -c -p 7000
```

### Removal of the cluster

```bash
docker-compose down
```

---

## Team

|                                                   David Procházka                                                   |                                                  Josef Podaný                                                  |                                                   Matěj Tužil                                                 |
| :-----------------------------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------------: | :---------: |
| [![David Procházka](https://avatars3.githubusercontent.com/u/2158418?s=200&v=4)](https://github.com/ProchazkaDavid) | [![Josef Podaný](https://avatars1.githubusercontent.com/u/30719925?s=200&v=4)](https://github.com/josefpodany) | [![Matěj Tužil](https://avatars.githubusercontent.com/xtuzil)](https://github.com/xtuzil)            |
|                          [`github.com/ProchazkaDavid`](https://github.com/ProchazkaDavid)                           |                           [`github.com/josefpodany`](https://github.com/josefpodany)                           |    [`github.com/xtuzil`](https://github.com/xtuzil)          |

---

## Sources
https://github.com/itsmetommy/docker-redis-cluster
https://github.com/Goodsmileduck/redis-cli
