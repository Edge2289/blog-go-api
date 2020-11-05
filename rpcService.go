package main

//import (
//	"blog-go-api/common"
//	pb "blog-go-api/proto"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"log"
//	"net"
//)
//
//const (
//	PORT = ":50001"
//)
//
//type server struct{}
//type additionServer struct {}
//
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//	log.Println("request: ", in.Name)
//	log.Println("request: ", in)
//
//	return &pb.HelloResponse{Message: "Hello 123415645" + common.MD5(in.Name)}, nil
//}
//
//func (s *additionServer) Addition(ctx context.Context, in *pb.AdditionRequest) (*pb.AdditionResponse, error) {
//	log.Println("Addition request: ", in.Name)
//	return &pb.AdditionResponse{Message:  in.Name + in.Name}, nil
//}
//
//func main() {
//	lis, err := net.Listen("tcp", PORT)
//
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//	pb.RegisterHelloServer(s, &server{})
//	pb.RegisterAdditionServer(s, &additionServer{})
//	log.Println("rpc服务已经开启")
//	s.Serve(lis)
//}
