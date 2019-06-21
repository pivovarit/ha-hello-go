FROM golang:1.12.6-alpine3.10 AS temp

ADD producer.go /producer.go

RUN go build -o /service /producer.go

FROM alpine:latest

COPY --from=temp /service .

EXPOSE 8080

CMD [ "./service" ]
