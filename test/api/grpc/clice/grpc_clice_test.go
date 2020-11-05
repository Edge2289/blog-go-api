package clice

import (
	pb "blog-go-api/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	address = "localhost:50001"
)


/**
获取搜索文章列表
*/
func TestArticleList(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewHelloClient(conn)
	//
	name := "lin ---- xxx"
	////if len(os.Args) > 1 {
	////	name = os.Args[1]
	////}
	//
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//
	fmt.Println("-----------r----------------")
	fmt.Println(r.Message)
	fmt.Println("------------r---------------")

	 var intName int64 = 10
	a := pb.NewAdditionClient(conn)
	ar, _ := a.Addition(context.Background(), &pb.AdditionRequest{Name: intName})
	fmt.Println("----------a-----------------")
	fmt.Println(ar)
	fmt.Println("-----------a----------------")
}