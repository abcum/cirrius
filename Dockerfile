FROM alpine:3.5

RUN apk update

RUN apk add --no-cache ca-certificates

ADD dep dep/

ADD cirrius .

EXPOSE 80 443

ENTRYPOINT ["/cirrius"]
