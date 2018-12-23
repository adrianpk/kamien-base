#!/bin/bash
# gif: generate - install - fresh
clear
killall -9 fresh
killall -9 runner-build
rm -rf views/views.go
go generate
go install
fresh
