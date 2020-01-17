FROM golang:1.11

COPY . /app
WORKDIR /app
RUN mkdir -p /app/bin/

RUN go mod download
RUN go build -o swagger-doc-server -v -ldflags '-s -w' main.go
RUN cp swagger-doc-server /app/bin/
