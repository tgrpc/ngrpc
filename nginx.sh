#!/bin/sh

sudo nginx -s stop
sudo nginx -c $GOPATH/src/github.com/tgrpc/ngrpc/nginx.conf
go run server.go >server1.log 2>&1 &
go run server2.go >server2.log 2>&1 &
tail -f /usr/local/nginx/logs/access.log
