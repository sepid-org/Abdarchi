FROM golang:1.22

ADD main.go /go/src/main.go

EXPOSE 8000

WORKDIR /go/src

ENTRYPOINT [ "go", "run",  "server.go"]