FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD gui gui/
ADD magnifio .

EXPOSE 80 443

ENTRYPOINT ["/magnifio"]