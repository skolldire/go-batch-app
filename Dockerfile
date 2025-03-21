ARG APP_NAME=go-batch-app

FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o $APP_NAME ./cmd/cli

RUN ls -la /app/$APP_NAME

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/$APP_NAME /$APP_NAME

RUN chmod +x /$APP_NAME

COPY --from=builder /app/config /app/config

ENV CONF_DIR=/app/config
ENV SCOPE=stage

RUN chmod -R 755 /app/config

CMD ["/$APP_NAME"]