package grpc_order

import (
	"context"
	"net"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	order_pb "github.com/Tomelin/fc-desafio-db/clean-architecture/pkg/grpc/order_pb"
)

type OrderHandlerGrpc struct {
	Service service.ServiceOrderInterface
	order_pb.UnimplementedOrderServiceServer
}

type HandlerGrpc interface {
	FindAll(context.Context, *order_pb.Empty) (*order_pb.ListOrderResponse, error)
	FindByFilter(context.Context, *order_pb.Filter) (*order_pb.ListOrderResponse, error)
	FindByID(context.Context, *order_pb.IdRequest) (*order_pb.OrderResponse, error)
}

// func NewOrderHandlerGrpc(svc service.ServiceOrderInterface) HandlerGrpc {
func NewOrderHandlerGrpc(svc service.ServiceOrderInterface) {

	order := &OrderHandlerGrpc{
		Service: svc,
	}

	err := order.Server(context.Background())
	if err != nil {
		panic(err)
	}
}

func (oh OrderHandlerGrpc) Server(ctx context.Context) error {
	grpcServer := grpc.NewServer()

	order_pb.RegisterOrderServiceServer(grpcServer, oh)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (oh OrderHandlerGrpc) FindAll(ctx context.Context, in *order_pb.Empty) (*order_pb.ListOrderResponse, error) {

	result, err := oh.Service.FindAll()
	if err != nil {
		return nil, err
	}

	var orders []*order_pb.OrderResponse
	for _, v := range result {
		o := &order_pb.OrderResponse{
			Id: v.ID,
			Order: &order_pb.Order{
				Name:        v.Order.Name,
				Description: v.Order.Description,
				Stock:       uint32(v.Order.Stock),
				Price:       v.Order.Price,
				Amount:      uint32(v.Order.Amount),
				Category:    1,
			},
		}

		orders = append(orders, o)
	}

	return &order_pb.ListOrderResponse{
		Orders: orders,
	}, nil
}

func (oh OrderHandlerGrpc) FindByFilter(ctx context.Context, in *order_pb.Filter) (*order_pb.ListOrderResponse, error) {
	result, err := oh.Service.FindByFilter(&in.Value)
	if err != nil {
		return nil, err
	}

	var orders []*order_pb.OrderResponse
	for _, v := range result {
		o := &order_pb.OrderResponse{
			Id: v.ID,
			Order: &order_pb.Order{
				Name:        v.Order.Name,
				Description: v.Order.Description,
				Stock:       uint32(v.Order.Stock),
				Price:       v.Order.Price,
				Amount:      uint32(v.Order.Amount),
				Category:    1,
			},
		}

		orders = append(orders, o)
	}

	return &order_pb.ListOrderResponse{
		Orders: orders,
	}, nil
}

func (oh OrderHandlerGrpc) FindByID(ctx context.Context, in *order_pb.IdRequest) (*order_pb.OrderResponse, error) {

	result, err := oh.Service.FindByID(&in.Id)
	if err != nil {
		return nil, err
	}

	return &order_pb.OrderResponse{
		Id: result.ID,
		Order: &order_pb.Order{
			Name:        result.Order.Name,
			Description: result.Order.Description,
			Stock:       uint32(result.Order.Stock),
			Price:       result.Order.Price,
			Amount:      uint32(result.Order.Amount),
			Category:    1,
		},
	},nil

}
