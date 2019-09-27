FROM alpine

EXPOSE  8086

VOLUME ["/data/templatego"]

RUN apk add --no-cache ca-certificates su-exec

COPY ./templatego /bin/templatego
ENV PUID=1000 PGID=1000

ENTRYPOINT \
  su-exec "${PUID}:${PGID}" \
     /bin/templatego \
       -c /data/templatego/config.yaml
