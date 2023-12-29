FROM golang:1.18 AS builder
WORKDIR /data
COPY . .
RUN go clean -modcache && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server .

FROM alpine:latest AS STANDARD

RUN mkdir -p /data/conf

WORKDIR /data
COPY --from=builder /data/server ./server
COPY --from=builder /data/conf/config.yml ./conf/config.yml

ENTRYPOINT ["/data/server"]
