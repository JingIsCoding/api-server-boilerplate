package response

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

func NewPagination(page int, pageSize int, total int64) Pagination {
	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}
