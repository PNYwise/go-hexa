package handlers

import "go-hexa/internal/core/domain/models/responses"

type ApiResponseList[T interface{}] struct {
	Message    string                        `json:"message"`
	Code       uint16                        `json:"code"`
	Data       T                             `json:"data"`
	Pagination *responses.PaginationResponse `json:"pagination"`
}

type ApiResponseDetail[T interface{}] struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
	Data    T      `json:"data"`
}

// NewApiResponseList is a function to create a new ApiResponseList instance
func NewApiResponseList[T interface{}](code uint16, data T, pagination *responses.PaginationResponse) *ApiResponseList[T] {
	return &ApiResponseList[T]{
		Message:    "success",
		Code:       code,
		Data:       data,
		Pagination: pagination,
	}
}

// NewApiResponseDetail is a function to create a new ApiResponseDetail instance
func NewApiResponseDetail[T interface{}](code uint16, data T) *ApiResponseDetail[T] {
	return &ApiResponseDetail[T]{
		Message: "success",
		Code:    code,
		Data:    data,
	}
}
