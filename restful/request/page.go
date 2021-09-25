package request

import "strconv"

const DefaultPageSize = 25
const MaxPageSize = 100

type PageRequest struct {
	Page     int
	PageSize int
}

func (request PageRequest) Validate() PageRequest {
	if request.Page < 0 {
		request.Page = 0
	}
	if request.PageSize <= 0 {
		request.PageSize = DefaultPageSize
	}
	if request.PageSize > MaxPageSize {
		request.PageSize = MaxPageSize
	}
	return request
}

func NewPageRequest(pageAsString string, pageSizeAsString string) PageRequest {
	var page, pageSize int64
	var err error
	if page, err = strconv.ParseInt(pageAsString, 10, 32); err != nil {
		page = 0
	}
	if pageSize, err = strconv.ParseInt(pageSizeAsString, 10, 32); err != nil {
		pageSize = DefaultPageSize
	}
	return PageRequest{
		Page:     int(page),
		PageSize: int(pageSize),
	}
}
