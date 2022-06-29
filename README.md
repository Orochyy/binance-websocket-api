go u# binance-websocket-api

###Clone app
```
https://github.com/Orochyy/binance-websocket-api.git
```

```
cd binance-websocket-api
```

###Build
```
pack build --builder gcr.io/buildpacks/builder:v1 binance-websocket-api
```

###Run it with docker, like:
```
docker run --rm -p 8080:8080 binance-websocket-api
```

###Run Locally with Go installed:
```
go build
./binance-websocket-api
```

###Run Locally with pack & Docker:
```
pack build --builder=gcr.io/buildpacks/builder binance-websocket-api
docker run -p8080:8080 binance-websocket-api
```

###Build Dockerfile (<binance-websocket-api> as name)
```
docker build . -t binance-websocket-api
```

###Run container
```
docker run -p 1010:1010 binance-websocket-api
```