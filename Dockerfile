FROM golang:1.15-alpine

WORKDIR /src

RUN apk add --no-cache make gcc musl-dev linux-headers git curl

ADD . .

ENTRYPOINT ["/src/docker-entrypoint.sh"]
