#!/bin/sh

go get golang.org/x/tools/cmd/godoc

# run go doc server
godoc -http=0.0.0.0:8080