package responses

type PaginationResponse struct {
	Page      int   `json:"page"`
	Take      int   `json:"take"`
	ItemCount int64 `json:"item_count"`
	PageCount int   `json:"page_count"`
}
