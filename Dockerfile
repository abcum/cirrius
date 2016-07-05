FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD app app/
ADD cirrius .

EXPOSE 80 443

ENTRYPOINT ["/cirrius"]