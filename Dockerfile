FROM golang:alpine

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

WORKDIR /go/icpscan
COPY . .

RUN go generate && go env && go build -o icpscan .

FROM alpine:latest
LABEL MAINTAINER="cwb2819259@gmail.com"

WORKDIR /go/icpscan

COPY --from=0 /go/icpscan ./

EXPOSE 8080

ENTRYPOINT ./icpscan
