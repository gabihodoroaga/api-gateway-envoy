FROM golang:1.17.11-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/bin/service1 -tags=jsoniter,nomsgpack ./service1

FROM scratch AS bin
WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/bin/service1 .

CMD ["/app/service1"]
