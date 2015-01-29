


var userRegex = regexp.MustCompile(`^(?P<user_id>\d+).\$cloud.(?P<device_id>\w+).channel.(?P<channel_id>[a-zA-Z0-9-_]+).event.state$`)


123.$device.b6b984190f.channel.on-off.event.state
456.$device.b6b984190f.channel.on-off.event.state

123.$device.b6b984190f.channel.on-off.event.state
123.$device.b6b984190f.channel.on-off.event.state



^(?P<user_id>\d+).\$cloud.(?P<device_id>\w+).channel.(?P<channel_id>[a-zA-Z0-9-_]+).event.state$


# Docker 

```
docker run --name ninja-rabbit -p 5672:5672 -p 15672:15672 -d mikaelhg/docker-rabbitmq
docker run --name ninja-redis -d redis
```