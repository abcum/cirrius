FROM alpine:latest

RUN apk add --update --no-cache ca-certificates

ADD dep dep/

ADD cirrius /usr/bin/

ENTRYPOINT ["cirrius"]
