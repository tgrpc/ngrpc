#!/bin/sh

sudo nginx -s stop
sudo nginx -c $GOPATH/src/github.com/tgrpc/ngrpc/nginx.conf
go run server.go &
tail -f /usr/local/nginx/logs/access.log
