FROM golang:1.13.4-alpine

RUN apk add --update --no-cache ca-certificates git

WORKDIR /src/

ENTRYPOINT ["sh","script/entrypoint.sh"]
