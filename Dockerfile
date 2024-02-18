FROM golang:1.22

ADD server.go /go/src/server.go

EXPOSE 8000

WORKDIR /go/src

ENTRYPOINT ["go", "run",  "server.go"]