package response

type PageData struct {
	Data     any   `json:"data"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total"`
}

func NewPageData(data any, page, pageSize, total int64) *PageData {
	return &PageData{
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}
