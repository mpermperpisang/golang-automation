# Golang Automation
Feature List :
- [x] API
- [x] Web
- [x] Android
- [x] iOS

## MODULE ENV
- [x] go env -w GO111MODULE=on

## REPOSITORY
- [x] Clone repo inside default $GOPATH (Example: `$HOME/go/src/github.com/`)
or
- [x] Clone repo outside default $GOPATH (Example: `$HOME/Documents/project/`)

## SPECS
- [x] macOS Big Sur 11.4
- [x] xcode 12.5.1
- [x] genymotion 3.2.1
- [x] java 1.8.0_181
- [x] gradle 7.2
- [x] carthage 0.38.0
- [x] cmake 3.13.1
- [x] node 14.15.4
- [x] npm 6.14.10
- [x] appium server 1.21.0
- [x] appium doctor 1.16.0
- [x] appium desktop 1.21.0
- [x] go 1.17
- [x] jenkins 2.289.3

## PROFILE
```
export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-1.8.0.jdk/Contents/Home
export ANDROID_HOME=$HOME/Library/Android/sdk
export PATH=$JAVA_HOME/bin:$PATH
export PATH=$PATH:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools
export GOROOT=/usr/local/go
export GOPATH=$HOME/Documents/project/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## MAKEFILE
- [x] `make cp`
- [x] `make deps`
- [x] `docker-compose up -d` (after test, run this command `docker-compose down`)

## RUN
- [x] <b>`make api-godog`</b>
- [x] <b>`make api-ginkgo`</b>
- [x] <b>`ginkgo -p --randomize-all --label-filter="get"`</b>. Need to `cd ./features/scenarios/non-xray/non-cucumber/api` first.

## CONTACT
`mpermperpisang@gmail.com`

## REFERENCES
- https://krishnachetan.medium.com/setup-appium-on-mac-1e06f1178427
- https://github.com/sayems/golang.webdriver
- https://medium.com/inside-bukalapak/recipe-to-boil-web-automation-with-go-language-98b715800d70
- https://mpermperpisang.medium.com/go-vatar-series-basic-automation-api-with-golang-f817efe217b5
- https://mpermperpisang.medium.com/go-vatar-series-android-and-golang-alliance-to-form-an-automation-bb5e25773802
- https://github.com/appium/appium/tree/master/sample-code/apps
- https://maelvls.dev/go111module-everywhere/
- https://github.com/eficode/Docker-Selenium-Example
