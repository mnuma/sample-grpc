package main

import (
	"google.golang.org/grpc"
	"log"
	"context"
	"fmt"

	pb "github.com/mnuma/sample-grpc/suga/proto"
)

func main() {
	 conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	 if err != nil {
			 log.Fatal("client connection error:", err)
	 }
	 defer conn.Close()
	 client := pb.NewEraClient(conn)
	 message := &pb.CurrentEraMessage{}
	 res, err := client.GetCurrentEra(context.TODO(), message)
	 fmt.Printf("新元号は『%s』 \n", res.Name)
}
