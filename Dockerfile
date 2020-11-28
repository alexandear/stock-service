FROM golang:1.15-alpine as build

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache git make

WORKDIR ./src/github.com/devchallenge/stock-service

COPY . ./

RUN make build && cp ./bin/stock-service /usr/local/bin/

FROM alpine

COPY --from=build /usr/local/bin/ /usr/local/bin/

ENV HOST 0.0.0.0

EXPOSE 9090

ENTRYPOINT ["stock-service"]
