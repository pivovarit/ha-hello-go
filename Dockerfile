FROM golang:1.12.6-alpine3.10

EXPOSE 8080

RUN mkdir /src
ADD hello.go /src/hello.go

RUN go build -o service /src/hello.go
ENTRYPOINT ./service
