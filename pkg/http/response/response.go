package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/ningzining/L-errors"
	log "github.com/ningzining/L-log"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResponse(code int, msg string, data any) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, NewResponse(CodeSuccess, MessageSuccess, data))
}

func Error(ctx *gin.Context, err error) {
	coder := errors.ParseCoder(err)
	if coder != nil {
		log.Error(coder.String())
		ctx.JSON(coder.HTTPStatus(), NewResponse(coder.Code(), coder.String(), nil))
	}
}
