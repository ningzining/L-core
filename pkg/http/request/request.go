package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context) uint64 {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	return id
}

func GetPageParam(ctx *gin.Context) *PageParam {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.Query("page_size"), 10, 64)

	return NewPageParam(page, pageSize)
}
