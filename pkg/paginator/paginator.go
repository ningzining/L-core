package paginator

import (
	"gorm.io/gorm"
)

type Paginator struct {
	Index int // 页码
	Size  int // 每页数量
}

func New(index, size int) *Paginator {
	if index <= 0 {
		index = 1
	}
	if size <= 0 {
		size = 10
	}
	return &Paginator{
		Index: index,
		Size:  size,
	}
}

func (p *Paginator) Offset() int {
	if p.Index <= 1 {
		return 0
	}
	return (p.Index - 1) * p.Size
}

func (p *Paginator) Interceptor() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset()).Limit(p.Size)
	}
}
