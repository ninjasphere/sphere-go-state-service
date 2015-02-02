# sphere-go-state-service

This service takes state messages from AMQP and stores them in REDIS.

# Docker 

```
docker run --name ninja-mysql -e MYSQL_ROOT_PASSWORD=ninja -d mysql
docker run --name ninja-rabbit -p 5672:5672 -p 15672:15672 -d mikaelhg/docker-rabbitmq
docker run --name ninja-redis -d redis
```

# Testing

```
123.$cloud.device.7511a8ecc5.channel.media.event.state

{"params":[{"media":{"id":"x-sonos-spotify:spotify%3atrack%3a3FUS56gKr9mVBmzvlnodlh?sid=12\u0026flags=32\u0026sn=1","type":"music-track","title":"Killing In The Name","duration":314000,"artists":[{"name":"Rage Against The Machine"}],"album":{"name":"Rage Against The Machine"}},"position":239000}],"jsonrpc":"2.0","time":1422501653158}

123.$cloud.device.7511a8ecc5.channel.volume.event.state

{"params":[{"level":0.19,"muted":false}],"jsonrpc":"2.0","time":1422501653233}

123.$cloud.device.7511a8ecc5.channel.media.event.state

{"params":[{"media":{"id":"x-sonos-spotify:spotify%3atrack%3a3FUS56gKr9mVBmzvlnodlh?sid=12\u0026flags=32\u0026sn=1","type":"music-track","title":"Killing In The Name","duration":314000,"artists":[{"name":"Rage Against The Machine"}],"album":{"name":"Rage Against The Machine"}},"position":239000}],"jsonrpc":"2.0","time":1422501658309}
```
