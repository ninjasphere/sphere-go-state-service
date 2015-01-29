


var userRegex = regexp.MustCompile(`^(?P<user_id>\w+).\$cloud.(?P<device_id>\w+).channel.(?P<channel_id>[a-zA-Z0-9-_]+).event.state$`)


123.$device.b6b984190f.channel.on-off.event.state
456.$device.b6b984190f.channel.on-off.event.state

123.$device.b6b984190f.channel.on-off.event.state
123.$device.b6b984190f.channel.on-off.event.state



^(?P<user_id>\d+).\$cloud.(?P<device_id>\w+).channel.(?P<channel_id>[a-zA-Z0-9-_]+).event.state$


# Docker 

```
docker run --name ninja-mysql -e MYSQL_ROOT_PASSWORD=ninja -d mysql
docker run --name ninja-rabbit -p 5672:5672 -p 15672:15672 -d mikaelhg/docker-rabbitmq
docker run --name ninja-redis -d redis
```


# Testing
123.$cloud.7511a8ecc5.channel.media.event.state
mosquitto_pub -t '123/$cloud/7511a8ecc5/channel/media/event/state' -h crusty.local -u guest -P guest -m '{"params":[{"media":{"id":"x-sonos-spotify:spotify%3atrack%3a3FUS56gKr9mVBmzvlnodlh?sid=12\u0026flags=32\u0026sn=1","type":"music-track","title":"Killing In The Name","duration":314000,"artists":[{"name":"Rage Against The Machine"}],"album":{"name":"Rage Against The Machine"}},"position":239000}],"jsonrpc":"2.0","time":1422501653158}'
EOF


Received PUBLISH (d0, q0, r0, m0, '$device/7511a8ecc5/channel/volume/event/state', ... (78 bytes))
{"params":[{"level":0.19,"muted":false}],"jsonrpc":"2.0","time":1422501653233}
Received PUBLISH (d0, q0, r0, m0, '$device/7511a8ecc5/channel/media/event/state', ... (335 bytes))
{"params":[{"media":{"id":"x-sonos-spotify:spotify%3atrack%3a3FUS56gKr9mVBmzvlnodlh?sid=12\u0026flags=32\u0026sn=1","type":"music-track","title":"Killing In The Name","duration":314000,"artists":[{"name":"Rage Against The Machine"}],"album":{"name":"Rage Against The Machine"}},"position":239000}],"jsonrpc":"2.0","time":1422501658309}
Received PUBLISH (d0, q0, r0, m0, '$device/7511a8ecc5/channel/volume/event/state', ... (78 bytes))
{"params":[{"level":0.19,"muted":false}],"jsonrpc":"2.0","time":1422501658371}
Received PUBLISH (d0, q0, r0, m0, '$device/7511a8ecc5/channel/media/event/state', ... (335 bytes))
{"params":[{"media":{"id":"x-sonos-spotify:spotify%3atrack%3a3FUS56gKr9mVBmzvlnodlh?sid=12\u0026flags=32\u0026sn=1","type":"music-track","title":"Killing In The Name","duration":314000,"artists":[{"name":"Rage Against The Machine"}],"album":{"name":"Rage Against The Machine"}},"position":242000}],"jsonrpc":"2.0","time":1422501660896}
Received PUBLISH (d0, q0, r0, m0, '$device/7511a8ecc5/channel/volume/event/state', ... (78 bytes))
{"params":[{"level":0.19,"muted":false}],"jsonrpc":"2.0","time":1422501661307}
Received PUBLISH (d0, q0, r0, m0, '$device/b6b984190f/channel/on-off/event/state', ... (54 bytes))
{"params":[true],"jsonrpc":"2.0","time":1422501661283}
Received PUBLISH (d0, q0, r0, m0, '$device/b6b984190f/channel/brightness/event/state', ... (67 bytes))
{"params":[0.996078431372549],"jsonrpc":"2.0","time":1422501661287}
Received PUBLISH (d0, q0, r0, m0, '$device/b6b984190f/channel/color/event/state', ... (91 bytes))
{"params":[{"mode":"temperature","temperature":2808}],"jsonrpc":"2.0","time":1422501661287}
Received PUBLISH (d0, q0, r0, m0, '$device/b6b984190f/channel/transition/event/state', ... (53 bytes))
{"params":[700],"jsonrpc":"2.0","time":1422501661288}
Received PUBLISH (d0, q0, r0, m0, '$device/1d5da0146e/channel/on-off/event/state', ... (54 bytes))
{"params":[true],"jsonrpc":"2.0","time":1422501661290}
Received PUBLISH (d0, q0, r0, m0, '$device/1d5da0146e/channel/brightness/event/state', ... (67 bytes))
{"params":[0.996078431372549],"jsonrpc":"2.0","time":1422501661290}
Received PUBLISH (d0, q0, r0, m0, '$device/1d5da0146e/channel/color/event/state', ... (91 bytes))
{"params":[{"mode":"temperature","temperature":2808}],"jsonrpc":"2.0","time":1422501661291}
Received PUBLISH (d0, q0, r0, m0, '$device/1d5da0146e/channel/transition/event/state', ... (53 bytes))
{"params":[700],"jsonrpc":"2.0","time":1422501661291}