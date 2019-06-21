FROM golang:1.12.6-alpine3.10 as temp

ADD hello.go /hello.go

RUN go build -o /hello /hello.go

FROM alpine:latest

COPY --from=temp /hello .

EXPOSE 8080

CMD [ "./hello" ]
