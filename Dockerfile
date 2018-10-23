FROM golang:1.11.1-alpine

WORKDIR /go/app
COPY ./app .

RUN apk update && \
    apk upgrade -U -a && \
    apk add git


RUN export GOPATH=$PWD
RUN go get -d -v ./...
CMD ["go","run","main.go"]