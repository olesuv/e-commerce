FROM golang:1.22.2-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git gcc musl-dev

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

# Run
CMD ["/app"]
