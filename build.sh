#!/bin/bash  

# if permission denied
# run script with ` chmod +x build.sh ` 
readonly ServerName="GoTradeBackServer"

# rm
# rm ./$ServerName.tar.gz ./sgrid_app

# compile
go build -o $ServerName

# build
# tar -cvf $ServerName.tar.gz ./sgrid.yml ./sgrid_app ./dist
