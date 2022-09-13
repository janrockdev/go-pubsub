# Go-PubSub (NATS)

IN-PROGRESS (initial CLI in place)

### Configuration:
```shell
# default > config.json
{
  "redis_url": "localhost:6379"
}
```

### Build:
```shell
go build ./cmd/cli 
```

### Run:
```shell
./cli redis ping
```

### Dependencies:
```shell
#REDIS - registry
<toadd>

#NATS
docker run -p 4222:4222 -ti nats:latest
```