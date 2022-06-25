FROM golang:1.17.11-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/bin/envtemplate -tags=jsoniter,nomsgpack ./envtemplate

FROM envoyproxy/envoy-alpine:v1.21.1

ENV PORT=8080

WORKDIR /app

COPY --from=build /src/bin/envtemplate envtemplate
COPY gateway/envoy-template.yaml envoy-template.yaml
COPY gateway/start-envoy.sh start-envoy.sh

RUN chmod -R a+rw /app
RUN chmod +x /app/envtemplate

EXPOSE 8080

ENTRYPOINT ["/app/start-envoy.sh"]
