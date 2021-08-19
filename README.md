# golang-automation

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
export JAVA_HOME=$(/usr/libexec/java_home)
export ANDROID_HOME=/Users/bukalapak/Library/Android/sdk
export PATH=$JAVA_HOME/bin:$PATH
export PATH=$PATH:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools
export PATH=$PATH:/Applications/Genymotion.app/Contents/MacOS/tools
export GOROOT=/usr/local/go
export GOPATH=$HOME/project/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## PACKAGE
- [x] go mod download

## ENV
- [x] cp env.sample .env

## COMMAND
- [x] godog --tags=@scenarios --random --format=cucumber > test/report/cucumber_report.json

## CONTACT
`mpermperpisang@gmail.com`

## REFERENCES
- https://krishnachetan.medium.com/setup-appium-on-mac-1e06f1178427
- https://github.com/sayems/golang.webdriver
- https://medium.com/inside-bukalapak/recipe-to-boil-web-automation-with-go-language-98b715800d70
- https://mpermperpisang.medium.com/go-vatar-series-basic-automation-api-with-golang-f817efe217b5
- https://mpermperpisang.medium.com/go-vatar-series-android-and-golang-alliance-to-form-an-automation-bb5e25773802
