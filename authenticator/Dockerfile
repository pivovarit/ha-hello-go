FROM golang:1.12.6-alpine3.10 as builder

ADD authenticator.go /authenticator.go

RUN go build -o /authenticator /authenticator.go

FROM alpine:latest

COPY --from=builder /authenticator .

EXPOSE 8080

CMD [ "./authenticator" ]
