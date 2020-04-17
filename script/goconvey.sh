#!/bin/sh
kill -9 $(lsof -i:8080 -t)
$GOPATH/bin/goconvey
