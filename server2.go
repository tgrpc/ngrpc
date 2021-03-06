package main

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/google/gops/agent"
	"github.com/soheilhy/cmux"
	"github.com/tgrpc/ngrpc/helloworld"
	"github.com/tgrpc/ngrpc/webapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("req:%+v\n", in)
	md := GetInMetadata(ctx)
	log.Printf("req metadate: %+v\n", md)
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloV2(req *helloworld.HelloRequest, gs helloworld.Greeter_SayHelloV2Server) error {
	log.Printf("req:%+v\n", req)
	for i := 0; i < 100; i++ {
		gs.Send(&helloworld.HelloReply{Message: "Hello " + strconv.Itoa(i)})
	}
	return nil
}

func GetInMetadata(ctx context.Context) metadata.MD {
	md, _ := metadata.FromIncomingContext(ctx)
	return md
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("...")
	log.Println(port)

	mux := cmux.New(lis)
	grpcl := mux.Match(cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	httpl := mux.Match(cmux.HTTP1Fast())

	go func() {
		s := grpc.NewServer()
		svr := &server{}
		helloworld.RegisterGreeterServer(s, svr)
		log.Println("grpc serve...")
		if err := s.Serve(grpcl); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		webapi.RegisteHelloService()
		webapiS := &http.Server{
			Handler: webapi.WebapiService,
		}
		log.Println("webapi serve...")
		if err := webapiS.Serve(httpl); err != cmux.ErrListenerClosed {
			panic(err)
		}
	}()

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	log.Println("cmux serve...")
	if err := mux.Serve(); err != nil {
		panic(err)
	}
}
