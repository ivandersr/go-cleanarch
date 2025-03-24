package service

import (
	"context"

	"github.com/ivandersr/cleanarch/internal/infra/grpc/pb"
	"github.com/ivandersr/cleanarch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetAllOrdersUseCase usecase.GetAllOrdersUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	getAllOrdersUseCase usecase.GetAllOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Empty) (*pb.ListOrdersResponse, error) {
	output, err := s.GetAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	response := &pb.ListOrdersResponse{}

	for _, order := range output.Orders {
		var parsedOrder pb.OrderResponse
		parsedOrder.Id = order.ID
		parsedOrder.Price = float32(order.Price)
		parsedOrder.Tax = float32(order.Tax)
		parsedOrder.FinalPrice = float32(order.FinalPrice)
		response.Orders = append(response.Orders, &parsedOrder)
	}

	return response, nil
}
