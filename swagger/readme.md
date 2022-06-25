# swagger

This project contains the unified swagger definition for all services:

## How to

Install go-swagger

```bash
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o /usr/local/bin/swagger -L'#' "$download_url"
chmod +x /usr/local/bin/swagger
```

Combine swagger definitions

```bash
swagger mixin swagger.host.json ../service1/docs/swagger.json ../service2/docs/swagger.json > swagger.json
```

Run locally 

```bash
docker build -t nginx-swag:alpine .

docker run -d \
    --name nginx-swag \
    --network config_traefik_default \
    -e SWAGGER_HOST_NAME="" \
    nginx-swag:alpine

docker rm -f nginx-swag
```
