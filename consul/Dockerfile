FROM consul:1.5.2

RUN apk add --no-cache bash curl
RUN test -e /etc/nsswitch.conf || echo 'hosts: files dns' > /etc/nsswitch.conf

CMD ["agent", "-dev", "-client", "0.0.0.0"]