
grpcurl -d '{"name":"ngrpc"}' -protoset helloworld.protoset localhost:80 helloworld.Greeter/SayHello