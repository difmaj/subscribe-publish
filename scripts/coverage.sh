#!/bin/sh
go test -coverprofile /tmp/cover.out $(go list ./...) 
go tool cover -html=/tmp/cover.out -o /tmp/cover.html
google-chrome /tmp/cover.html