#!/bin/bash
# mk: make keys
openssl genrsa -out ./resources/keys/{{.AppNameLowercase}}.rsa 1024
openssl rsa -in ./resources/keys/{{.AppNameLowercase}}.rsa -pubout > ./resources/keys/{{.AppNameLowercase}}.rsa.pub