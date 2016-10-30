FROM golang:1.7
MAINTAINER Allen Zou <zyl_work@163.com>

RUN mkdir -p /go/libs
ENV GOPATH /go/libs

RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/go-sql-driver/mysql
RUN go get -v github.com/satori/go.uuid
RUN go get -v gopkg.in/yaml.v2
RUN go get -v gopkg.in/redis.v4
RUN go get -v github.com/Sirupsen/logrus

RUN mkdir -p /project
COPY ./go/src /project/src
ENV GOPATH $GOPATH:/project

RUN go install main

WORKDIR /project/bin
CMD ["./main"]
EXPOSE 80
