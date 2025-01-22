package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersOutputItemDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}
type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputItemDTO, error) {
	res, err := c.OrderRepository.ListAll()
	if err != nil {
		return []ListOrdersOutputItemDTO{}, err
	}

	dtos := []ListOrdersOutputItemDTO{}

	for _, order := range res {
		dto := ListOrdersOutputItemDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
