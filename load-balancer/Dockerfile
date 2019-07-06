FROM haproxy:2.0.1-alpine

RUN set -x	                                  &&  \
	apk upgrade --update					  &&  \
	apk add bash ca-certificates rsyslog	  &&  \
	mkdir -p /etc/rsyslog.d/				  &&  \
	touch /var/log/haproxy.log				  &&  \
	ln -sf /dev/stdout /var/log/haproxy.log

COPY etc/haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg

ADD ./etc/rsyslog.conf /etc/rsyslog.conf
RUN set -o errexit && set -o nounset && readonly RSYSLOG_PID="/var/run/rsyslogd.pid"
RUN rm -f $RSYSLOG_PID

COPY docker-entrypoint.sh docker-entrypoint.sh
RUN chmod +x docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]