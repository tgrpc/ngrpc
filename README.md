# ngrpc

Nginx-gRpc
============


 - 1 `nginx.conf`

[nginx-grpc-doc](https://www.nginx.com/blog/nginx-1-13-10-grpc/)

```
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"';

    server {
        listen 80 http2;
 
        access_log logs/access.log main;
 
        events {
          worker_connections  1024;  ## Default: 1024
        }

        location / {
            # Replace localhost:50051 with the address and port of your gRPC server
            # The 'grpc://' prefix is optional; unencrypted gRPC is the default    
            grpc_pass grpc://localhost:50051;
        }
    }
}

events {
    worker_connections  1024;  ## Default: 1024
}
```

 - 2 `server.go`

```
lis, err := net.Listen("tcp", ":50051")
```

 - 3 `client.go`

```
 conn, err := grpc.Dial("localhost:80", grpc.WithInsecure())
```