package request

type PageParam struct {
	PageIndex int64 `form:"page_index"` // 页码
	PageSize  int64 `form:"page_size"`  // 页大小
}

func NewPageParam(pageIndex, pageSize int64) *PageParam {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return &PageParam{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
}

func (p PageParam) GetPage() int64 {
	if p.PageIndex <= 0 {
		return 1
	}
	return p.PageIndex
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
