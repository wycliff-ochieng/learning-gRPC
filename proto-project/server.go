package main

import (
	pb "github.com/wycliff-ochieng/proto-project/coffee_proto"
	"context"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *Server) GetMenu(menuRequest *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := *[]pb.Items{
		&pb.Items{"id":"2","name":"Strong Tea"},
		&pb.Items{"id":"3","name":"Mursik"},
		&pb.Items{"id":"4","name":"white Tea"},
	}

	for i,_ := range items{
		srv.Send(&pb.Menu{
			Items:items[0 : i+1],
		})
	}
	return nil
}

func (s *Server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return nil, nil
}
func (s *Server) GetOrderStatus(context.Context, *pb.Receipt) (*pb.OrderStatus, error) {
	return nil, nil
}