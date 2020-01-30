FROM golang:1.13-alpine

RUN apk add gcc libc-dev && mkdir /app

WORKDIR /app
CMD sh
