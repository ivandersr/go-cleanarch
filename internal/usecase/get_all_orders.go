package usecase

import (
	"github.com/ivandersr/cleanarch/internal/entity"
)

type GetAllOrdersOutputDTO struct {
	Orders []GetOrderOutputDTO `json:"orders"`
}

type GetOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{OrderRepository: OrderRepository}
}

func (c *GetAllOrdersUseCase) Execute() (GetAllOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return GetAllOrdersOutputDTO{}, err
	}
	ordersOutput := []GetOrderOutputDTO{}
	for _, order := range orders {
		ordersOutput = append(ordersOutput, GetOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return GetAllOrdersOutputDTO{Orders: ordersOutput}, nil
}
