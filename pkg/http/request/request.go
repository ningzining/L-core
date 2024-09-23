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
	pageIndex, _ := strconv.ParseInt(ctx.Query("page_index"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.Query("page_size"), 10, 64)

	return NewPageParam(pageIndex, pageSize)
}
