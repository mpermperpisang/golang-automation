FROM golang:latest
WORKDIR /go/src/github.com/golang-automation
RUN GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.10.0 && \
    go get -t -d github.com/tebeka/selenium && \
    go get github.com/sclevine/agouti && \
    go get github.com/joho/godotenv && \
    go get -u github.com/logrusorgru/aurora && \
    go get -u github.com/magiconair/properties && \
    go get github.com/yalp/jsonpath && \
    go get -u golang.org/x/lint/golint && \
    go get -u github.com/goccy/go-yaml && \
    go get github.com/brianvoe/gofakeit && \
    go get -u github.com/tidwall/gjson && \
    go get github.com/golang/mock/gomock && \
    go get github.com/golang/mock/mockgen && \
    go get github.com/stretchr/testify

FROM node:latest
WORKDIR /go/src/github.com/golang-automation
RUN npm install cucumber-html-reporter

FROM ruby:2.5.5
WORKDIR /go/src/github.com/golang-automation
COPY Gemfile /
RUN gem update --system && \
    gem install bundler && \
    bundle install
