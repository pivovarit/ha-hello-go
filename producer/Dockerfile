FROM golang:1.12.6-alpine3.10 AS temp

RUN mkdir /src
ADD producer.go /src/producer.go

RUN go build -o /src/service /src/producer.go

FROM alpine:latest

COPY --from=temp /src/service .

EXPOSE 8080

CMD [ "./service" ]
