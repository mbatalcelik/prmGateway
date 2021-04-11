# Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go get -d -v ./src
RUN go build -o demo-app ./src
# Stage 2
FROM alpine as alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./demo-app"]
# Stage 3
FROM registry.access.redhat.com/ubi8/go-toolset as builderRedhat
COPY --chown=default:root . /build/
WORKDIR /build
RUN go build -o demo-app
# Stage 4
FROM registry.access.redhat.com/ubi8/ubi-minimal as redhat
COPY --from=builderRedhat /build/ /app/
WORKDIR /app
CMD ["./demo-app"]