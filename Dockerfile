FROM golang:latest

WORKDIR /go/icpscan
COPY . .

RUN go generate && go env && go build -o icpscan .

FROM alpine:latest
LABEL MAINTAINER="cwb2819259@gmail.com"

WORKDIR /go/icpscan

COPY --from=0 /go/icpscan ./

EXPOSE 8080

ENTRYPOINT ./icpscan -c icpscan.config