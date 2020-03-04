# ngrpc


1.[gRPC和http使用相同端口](cmux.md)
============

2.Nginx-gRPC gRPC服务代理
============


 - 1 `nginx.conf` 开启nginx grpc代理

```
sudo nginx -s stop
sudo nginx -c $GOPATH/src/github.com/tgrpc/ngrpc/nginx.conf
tail -f /usr/local/nginx/logs/access.log
```

[nginx-grpc-doc](https://www.nginx.com/blog/nginx-1-13-10-grpc/)

```
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"';

    server {
        listen 2080 http2;
 
        access_log logs/access.log main;

        location /helloworld.Greeter {
            # Replace localhost:50052 with the address and port of your gRPC server
            # The 'grpc://' prefix is optional; unencrypted gRPC is the default    
            grpc_pass grpc://localhost:50052;
        }

        location /helloworld.LangService {
            # Replace localhost:50051 with the address and port of your gRPC server
            # The 'grpc://' prefix is optional; unencrypted gRPC is the default    
            grpc_pass grpc://localhost:50051;
        }
    }

    server {
        listen       80;
        server_name  localhost;
        client_max_body_size 1024M;

        location /api/helloworld.Greeter {
            proxy_pass http://127.0.0.1:50052;
            proxy_set_header Host $host:$server_port;
        }
    }
}

events {
    worker_connections  1024;  ## Default: 1024
}
```

 - 2 开启服务 `go run server.go`

```
lis, err := net.Listen("tcp", ":50051")
```

 - 3 测试服务 `go run client.go`

```
 conn, err := grpc.Dial("localhost:2080", grpc.WithInsecure())
```

## extend option

```
protoc -I . --gogo_out=plugins=grpc:extension extension/extension.proto
protoc -I . --gogo_out=plugins=grpc:. helloworld/greeter.proto
go run option/main.go
```

