FROM alpine:latest

RUN apk add --update --no-cache ca-certificates

ADD cirrius /usr/bin/

ENTRYPOINT ["cirrius"]
