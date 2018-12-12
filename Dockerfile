FROM alpine

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*

RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY bodyguard ./

ENV LOGGING_LEVEL WARNING

COPY docker-entrypoint.sh ./
ENTRYPOINT ["./docker-entrypoint.sh"]

CMD ["./bodyguard"]
