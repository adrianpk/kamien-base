#!/bin/bash
# mc: make certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout resources/certificates/key.pem -out resources/certificates/cert.pem
