package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wycliff-ochieng/proto-project/coffee_proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *Server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Items{
		&pb.Items{Id: "2", Name: "Strong Tea"},
		&pb.Items{Id: "3", Name: "Mursik"},
		&pb.Items{Id: "4", Name: "white Tea"},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{
			Item: items[0 : i+1],
		})
	}
	return nil
}

func (s *Server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Item: []*pb.Items{
			{Id: "1", Name: "strong tea"},
		},
	}, nil
}
func (s *Server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	var orderId string
	return &pb.OrderStatus{
		OrderId: orderId,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &Server{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %s", err)
	}

}
