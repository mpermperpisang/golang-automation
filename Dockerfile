FROM golang:1.14

WORKDIR /go/src/github.com/golang-automation
COPY . .

RUN make package
RUN make docker
