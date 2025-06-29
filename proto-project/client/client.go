package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wycliff-ochieng/proto-project/coffee_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect")
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	menustream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("")
	}

	done := make(chan bool)

	var items []*pb.Items

	go func() {
		for {
			resp, err := menustream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("")
			}
			items = resp.Item
			log.Printf("Resp is receiced : %v", resp.Item)
		}
	}()

	<-done
	log.Printf("channels is done %v", items)
}
