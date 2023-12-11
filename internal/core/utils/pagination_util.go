package utils

import (
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"math"
)

func GeneratePagination(paginationRequest *requests.PaginationRequest, count int64) *responses.PaginationResponse {
	totalPages := int(math.Ceil(float64(count) / float64(paginationRequest.Take)))
	pagination := new(responses.PaginationResponse)
	pagination.Page = paginationRequest.Page
	pagination.Take = paginationRequest.Take
	pagination.ItemCount = count
	pagination.PageCount = totalPages

	return pagination
}
