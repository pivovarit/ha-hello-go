FROM golang:1.12.6-alpine3.10 AS temp

ADD client.go /client.go

RUN go build -o /service /client.go

FROM alpine:latest

COPY --from=temp /service .

EXPOSE 8080

CMD [ "./service" ]
