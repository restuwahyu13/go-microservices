package main

import (
	"context"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/restuwahyu13/go-microservices/internals/schemas/orders"
)

type ordersService struct {
	orders.UnimplementedOrdersServiceServer
}

func main() {
	ls, _ := net.Listen("tcp", ":4001")
	server := grpc.NewServer()

	orders.RegisterOrdersServiceServer(server, &ordersService{})
	server.Serve(ls)
}

func (h *ordersService) PingOrders(ctx context.Context, in *emptypb.Empty) (*orders.GrpcResponse, error) {
	res := new(orders.GrpcResponse)
	res.StatCode = http.StatusOK
	res.StatMessage = "Orders Service Pong"

	return res, nil
}
