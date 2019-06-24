FROM golang:1.12.6-alpine3.10 as temp

ADD users.go /users.go

RUN go build -o /users /users.go

FROM alpine:latest

COPY --from=temp /users .

EXPOSE 8080

CMD [ "./users" ]
