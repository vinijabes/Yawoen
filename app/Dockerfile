FROM golang:latest

RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/mux

RUN mkdir /app

ADD . /app/

WORKDIR /app
RUN go build
ENTRYPOINT ./app

EXPOSE 5000