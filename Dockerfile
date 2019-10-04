FROM redis:latest

ARG port

RUN mkdir -p /etc/redis && \
    echo "port $port\n\
    cluster-enabled yes\n\
    cluster-config-file nodes.conf\n\
    cluster-node-timeout 5000\n\
    appendonly yes" > /etc/redis/redis.conf

EXPOSE $port

ENTRYPOINT ["redis-server", "/etc/redis/redis.conf"]