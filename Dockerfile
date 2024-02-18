FROM golang:1.22

ADD . /go/src/

EXPOSE 8000

WORKDIR /go/src

ENTRYPOINT ["go", "run",  "server.go"]