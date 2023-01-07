package main

import (
	"context"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/restuwahyu13/go-microservices/internals/schemas/products"
)

type productsService struct {
	products.UnimplementedProductsServiceServer
}

func main() {
	ls, _ := net.Listen("tcp", ":4002")
	server := grpc.NewServer()

	products.RegisterProductsServiceServer(server, &productsService{})
	server.Serve(ls)
}

func (h *productsService) PingProducts(ctx context.Context, in *emptypb.Empty) (*products.GrpcResponse, error) {
	res := new(products.GrpcResponse)
	res.StatCode = http.StatusOK
	res.StatMessage = "Products Service Pong"

	return res, nil
}
