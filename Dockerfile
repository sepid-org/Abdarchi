FROM golang:1.22

ADD . /go/src/

WORKDIR /go/src/

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server-go

EXPOSE 8000

CMD ["/server-go"]