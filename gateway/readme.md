# gateway

Simple envoy based api gateway

## How to

Build

```bash
docker build -t api-gateway .
```

Run local
```bash
docker run -d \
    --name api-gateway \
    -e PORT=8080 \
    -e SERVICE1_HOST=localhost \
    -e SERVICE1_PORT=8081 \
    -e SERVICE1_SSL=false \
    -e SERVICE2_HOST=localhost \
    -e SERVICE2_PORT=8082 \
    -e SERVICE2_SSL=false \
    -e SWAGGER_HOST=localhost \
    -e SWAGGER_PORT=8083 \
    -e SWAGGER_SSL=false \
    api-gateway
```

