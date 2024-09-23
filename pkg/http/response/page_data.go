package response

type PageData struct {
	Data      any   `json:"data"`
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
	Total     int64 `json:"total"`
}

func NewPageData(data any, pageIndex, pageSize, total int64) *PageData {
	return &PageData{
		Data:      data,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     total,
	}
}
