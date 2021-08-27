#!/bin/bash
host=127.0.0.1
port=6379
pattern='goods:modify:*'

redis-cli -h $host -p $port --scan --pattern "$pattern" | while read LINE ; do
    TTL=`redis-cli -h "$host" -p "$port" ttl "$LINE"`;
    if [ $TTL -eq  -1 ]; then
        echo "$LINE";
        redis-cli -h $host -p $port expire "$LINE" 15552000
    fi;
done;
