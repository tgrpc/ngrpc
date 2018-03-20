#!/bin/sh

sudo nginx -s reload
# go run server.go &
tail -f /usr/local/nginx/logs/access.log
