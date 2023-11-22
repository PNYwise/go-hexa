package requests

import "go-hexa/internal/core/domain/enums/order"

type PaginationRequest struct {
	Page  int        `json:"page" validate:"required,min=1"`
	Take  int        `json:"take" validate:"required,min=1"`
	Order order.Enum `json:"order" validate:"required,alpha"`
}

func (p *PaginationRequest) Skip() int {
	return (p.Page - 1) * p.Take
}
