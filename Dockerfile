
# STAGE 1
FROM golang:1.20-alpine AS build

ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN apk add curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN go build -o /cint-test

# STAGE 2
FROM alpine:3.17

ENV SERVE_PORT 8080
ENV GIN_MODE release

RUN apk update && apk add make "libcrypto3>=3.0.8-r1" "libssl3>=3.0.8-r1" && rm -rf /var/cache/apt/*

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /cint-test /bin/cint-test

USER 65534

ENTRYPOINT ["/bin/cint-test"]