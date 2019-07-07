FROM golang:1.12.6-alpine3.10 AS builder

ADD client.go /client.go

RUN go build -o /service /client.go

FROM alpine:latest

COPY --from=builder /service .

EXPOSE 8080

CMD [ "./service" ]
