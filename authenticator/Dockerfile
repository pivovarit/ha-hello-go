FROM golang:1.12.6-alpine3.10 as temp

ADD authenticator.go /authenticator.go

RUN go build -o /authenticator /authenticator.go

FROM alpine:latest

COPY --from=temp /authenticator .

EXPOSE 8080

CMD [ "./authenticator" ]
