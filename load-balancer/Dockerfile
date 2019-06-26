FROM haproxy:1.9.2-alpine

RUN set -x	&& \
	apk upgrade --update					&&  \
	apk add bash ca-certificates rsyslog	                &&  \
	mkdir -p /etc/rsyslog.d/				&&  \
	touch /var/log/haproxy.log				&&  \
	ln -sf /dev/stdout /var/log/haproxy.log

COPY etc/haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg

ADD ./etc/rsyslog.conf /etc/rsyslog.conf
RUN set -o errexit
RUN set -o nounset
RUN readonly RSYSLOG_PID="/var/run/rsyslogd.pid"
RUN rm -f $RSYSLOG_PID

COPY docker-entrypoint.sh docker-entrypoint.sh
RUN chmod +x docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["haproxy", "-f", "/usr/local/etc/haproxy/haproxy.cfg"]