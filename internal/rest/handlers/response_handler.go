package handlers

type ApiResponseList struct {
	Message    string      `json:"message"`
	Code       uint16      `json:"code"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type ApiResponseDetail struct {
	Message string      `json:"message"`
	Code    uint16      `json:"code"`
	Data    interface{} `json:"data"`
}

// NewApiResponseList is a function to create a new ApiResponseList instance
func NewApiResponseList(code uint16, data interface{}, pagination interface{}) *ApiResponseList {
	return &ApiResponseList{
		Message:    "success",
		Code:       code,
		Data:       data,
		Pagination: pagination,
	}
}

// NewApiResponseDetail is a function to create a new ApiResponseDetail instance
func NewApiResponseDetail(code uint16, data interface{}) *ApiResponseDetail {
	return &ApiResponseDetail{
		Message: "success",
		Code:    code,
		Data:    data,
	}
}
