FROM golang:1.22.2-alpine AS builder

RUN apk --no-cache add bash git gcc musl-dev

WORKDIR /usr/local/src

# Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Build
COPY . .
RUN go build -o ./bin/app ./server.go

# Clean alpine image
FROM alpine AS app

COPY --from=builder /usr/local/src/bin/app /
COPY .env .env

RUN apk --no-cache add redis

# Run
CMD ["/app", "redis-server"]
