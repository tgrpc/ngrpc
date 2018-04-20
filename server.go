/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"log"
	"net"
	"time"

	"github.com/tgrpc/ngrpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("context:%+v\n", ctx)
	log.Printf("req:%+v\n", in)
	md := GetInMetadata(ctx)
	log.Printf("req metadate: %+v\n", md)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) List(ctx context.Context, in *helloworld.ListReq) (*helloworld.ListResp, error) {
	return &helloworld.ListResp{
		Langs: []*helloworld.Lang{
			&helloworld.Lang{
				Name:     "Golang",
				Birthday: time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local).Unix(),
				Versions: []*helloworld.Version{
					&helloworld.Version{
						Id:      1,
						Version: "v0.0.1",
						Desc:    "go version 0.0.1",
						Vi32S:   []int32{111, 222},
						Vi64S:   []int64{111, 222},
						Vf64S:   []float64{111.0, 222.0},
						Vstrs:   []string{"str1", "str2"},
					},
					&helloworld.Version{
						Id:      190,
						Version: "v1.9.0",
						Desc:    "go version 1.9.0",
						Vi32S:   []int32{111, 222},
						Vi64S:   []int64{111, 222},
						Vf64S:   []float64{111.0, 222.0},
						Vstrs:   []string{"str1", "str2"},
					},
				},
			},
			&helloworld.Lang{
				Name:     "Java",
				Birthday: time.Date(1992, 1, 2, 15, 4, 5, 0, time.Local).Unix(),
				Versions: []*helloworld.Version{
					&helloworld.Version{
						Id:      700,
						Version: "v7.0.0",
						Desc:    "jdk7",
						Vi32S:   []int32{111, 222},
						Vi64S:   []int64{111, 222},
						Vf64S:   []float64{111.0, 222.0},
						Vstrs:   []string{"str1", "str2"},
					},
					&helloworld.Version{
						Id:      900,
						Version: "v9.0.0",
						Desc:    "jdk9",
						Vi32S:   []int32{111, 222},
						Vi64S:   []int64{111, 222},
						Vf64S:   []float64{111.0, 222.0},
						Vstrs:   []string{"str1", "str2"},
					},
				},
			},
		},
		TotalCount: 999,
	}, nil
}

func (s *server) GetLang(ctx context.Context, in *helloworld.GetLangReq) (*helloworld.Lang, error) {
	return &helloworld.Lang{
		Name:     "Golang",
		Birthday: time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local).Unix(),
		Versions: []*helloworld.Version{
			&helloworld.Version{
				Id:      1,
				Version: "v0.0.1",
				Desc:    "go version 0.0.1",
				Vi32S:   []int32{111, 222},
				Vi64S:   []int64{111, 222},
				Vf64S:   []float64{111.0, 222.0},
				Vstrs:   []string{"str1", "str2"},
			},
			&helloworld.Version{
				Id:      190,
				Version: "v1.9.0",
				Desc:    "go version 1.9.0",
				Vi32S:   []int32{111, 222},
				Vi64S:   []int64{111, 222},
				Vf64S:   []float64{111.0, 222.0},
				Vstrs:   []string{"str1", "str2"},
			},
		},
	}, nil
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
	s := grpc.NewServer()
	svr := &server{}
	pb.RegisterGreeterServer(s, svr)
	helloworld.RegisterLangServiceServer(s, svr)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println(port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
