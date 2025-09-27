FROM golang:1.19 AS builder

ADD . /repo
WORKDIR /repo

RUN go build -o bin/example

FROM debian:bookworm-slim
COPY --from=builder /repo/bin/example /usr/bin/example

ENTRYPOINT ./usr/bin/example