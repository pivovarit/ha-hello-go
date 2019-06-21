FROM golang:latest
RUN mkdir /src
ADD hello.go /src/hello.go

RUN go build -o service /src/hello.go
ENTRYPOINT ./service
