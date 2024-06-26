FROM golang:1.22-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install github.com/swaggo/swag/cmd/swag@latest && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
EXPOSE 8000
EXPOSE 2345

ENTRYPOINT ["air", "-c", ".air.toml"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o carsearch -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/my-great-program /usr/local/bin/carsearch
EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/carsearch"]