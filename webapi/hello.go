package webapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/tgrpc/ngrpc/helloworld"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	"encoding/json"
	"github.com/tgrpc/grpcurl"
	"google.golang.org/grpc/keepalive"
	"time"
)

var WebapiService *mux.Router

func init() {
	WebapiService = mux.NewRouter()
}

func SayHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("say hello,%+v \n", req)
	fmt.Fprintf(rw, "hello %+v", req)
}

func NewGreterClient() helloworld.GreeterClient {
	conn, err := grpc.Dial("127.0.0.1:2080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := helloworld.NewGreeterClient(conn)
	return c
}

func NewGreterClientV2() helloworld.GreeterClient {
	cc, err := grpcurl.BlockingDial(context.Background(), "tcp", "127.0.0.1:2080", nil, grpc.WithKeepaliveParams(
		keepalive.ClientParameters{
			Time:    time.Second * 10,
			Timeout: time.Second * 10,
		},
	))
	if err != nil {
		fmt.Printf("err:%+v", err)
		return nil
	}
	return helloworld.NewGreeterClient(cc)
}

func GrpcSayHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("GrpcSayHello...")
	um := new(jsonpb.Unmarshaler)
	pb := new(helloworld.HelloRequest)
	if err := um.Unmarshal(req.Body, pb); err != nil {
		fmt.Printf("err:%+v", err)
		return
	}

	fmt.Println("GrpcSayHello... call")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client := NewGreterClient()
	resp, err := client.SayHello(ctx, pb)
	// defer client.Close()
	if err != nil {
		fmt.Printf("err:%+v", err)
		return
	}
	HttpResEncode(rw, resp)
}

func HttpResEncode(rw http.ResponseWriter, resp interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	if msg, ok := resp.(proto.Message); ok {
		(&jsonpb.Marshaler{EmitDefaults: true}).Marshal(rw, msg)
	} else {
		ret, _ := json.Marshal(resp)
		rw.Write(ret)
	}
}

func RegisteHelloService() {
	// WebapiService.HandleFunc("/api/helloworld.Greeter/SayHello", SayHello)
	WebapiService.HandleFunc("/api/helloworld.Greeter/SayHello", GrpcSayHello)
}
