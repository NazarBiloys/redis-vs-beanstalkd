FROM golang:1.18-alpine as builder

WORKDIR /app

# Fetch dependencies.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary.
RUN cd cmd/redis-vs-beanstalkd \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/redis-vs-beanstalkd .

FROM alpine:3.16

ARG ROOT_FOLDER="app"
ARG BUILD_ID=0
ARG VERSION=0.0.1

ENV ROOT_FOLDER=${ROOT_FOLDER} \
    BUILD_ID=${BUILD_ID} \
    VERSION=${VERSION}

# Copy static executable
COPY --from=builder /go/bin/redis-vs-beanstalkd app/redis-vs-beanstalkd

EXPOSE 90

# Run binary.
ENTRYPOINT ["app/redis-vs-beanstalkd"]
