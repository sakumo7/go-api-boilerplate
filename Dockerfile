FROM golang:1.11.1-alpine

WORKDIR src/
COPY ./src .

RUN apk update && \
    apk upgrade -U -a && \
    apk add git

RUN go get github.com/gorilla/mux
CMD ["go","run","main.go"]