FROM golang:latest

WORKDIR /go/api-go
COPY . .

# RUN echo "teste"

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["echo teste"]