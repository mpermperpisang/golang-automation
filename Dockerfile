FROM golang:1.17
LABEL owner="Ferawati Hartanti Pratiwi" email=mpermperpisang@gmail.com

ENV GO111MODULE=auto
WORKDIR $HOME/Documents/project/golang-automation-v1
COPY Makefile go.mod go.sum *.sample ./
RUN make selenium && make deps && make cp

FROM node:14.15.4
