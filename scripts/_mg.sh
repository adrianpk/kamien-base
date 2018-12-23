#!/bin/sh
# mg: migration (create)
clear
migrate create -ext .sql -dir ./resources/migrations -format "20060102150405" $1