#!/bin/bash
# gi: generate - install
clear
rm -rf views/views.go
go generate
go install
