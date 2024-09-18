#!/bin/bash

# if permission denied
# run script with ` chmod +x build.sh `
readonly ServerName="GoTradeBackServer"

# rm
rm ./$ServerName.tar.gz ./$ServerName

# compile
# GOOS=linux GOARCH=amd64
go build -o $ServerName

# build
tar -cvf $ServerName.tar.gz  $ServerName

#  GOOS=linux GOARCH=amd64 go build -o GoTradeBackServer
