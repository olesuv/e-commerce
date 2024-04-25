FROM golang:alpine AS builder

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

RUN apk --no-cache add redis bash

COPY --from=builder /usr/local/src/bin/app /
COPY .env .env

# Run
CMD ["bash", "-c", "redis-server & ./app"]
