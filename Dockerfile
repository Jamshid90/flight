FROM golang:1.17 as builder

RUN GOCACHE=OFF

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN make build-linux

FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY --from=builder /app/bin/flight ./flight

RUN chmod +x ./flight

CMD ["./flight"]