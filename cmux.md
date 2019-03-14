## example

使用[github.com/soheilhy/cmux](https://github.com/soheilhy/cmux) 监听gRPC和http服务；
http接口实现：调用gRPC接口。

### code

1. 运行服务
```
go run server.go
```
2. gRPC调用
```
tg -c tgrpc.toml -C
```

3. http调用
```
curl http://localhost:50052/api/helloworld.Greeter/SayHello -H 'customerId:123' -H 'region:UK' --data-binary '{"name":"gRPC&webapi"}' --compressed
curl http://localhost:2080/api/helloworld.Greeter/SayHello -H 'customerId:123' -H 'region:UK' --data-binary '{"name":"gRPC&webapi"}' --compressed
```
