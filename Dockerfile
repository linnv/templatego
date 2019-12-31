FROM alpine

EXPOSE  8086

VOLUME ["/data/qnmock"]

RUN apk add --no-cache ca-certificates su-exec

COPY ./qnmock /bin/qnmock
ENV PUID=1000 PGID=1000

ENTRYPOINT \
  su-exec "${PUID}:${PGID}" \
     /bin/qnmock \
       -c /data/qnmock/config.yaml
