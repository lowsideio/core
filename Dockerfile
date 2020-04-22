# Dockerfile

FROM golang:1.9-alpine

MAINTAINER Aksels Ledins

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /usr/src/app

ADD . /usr/src/app

COPY . .

# dependencies
RUN make deps
RUN make build

EXPOSE 1323

ENTRYPOINT [ "/usr/src/app/core-api" ]
