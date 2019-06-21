FROM golang:1.12.6-stretch

RUN mkdir /src
ADD producer.go /src/producer.go

RUN go build -o service /src/producer.go
ENTRYPOINT ./service
