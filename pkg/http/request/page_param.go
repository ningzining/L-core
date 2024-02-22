package request

type PageParam struct {
	Page     int64 `form:"page"`      // 页码
	PageSize int64 `form:"page_size"` // 页大小
}

func NewPageParam(page, pageSize int64) *PageParam {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return &PageParam{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p PageParam) GetPage() int64 {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p PageParam) GetPageSize() int64 {
	if p.PageSize <= 0 {
		return 10
	}

	return p.PageSize
}

func (p PageParam) GetOffset() int64 {
	return p.GetPage() * p.GetPageSize()
}
