package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Data struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}

type ListResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Data  `json:"data"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type Response struct{}

func (r *Response) SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

func (r *Response) ErrorResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusInternalServerError, ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func (r *Response) SuccessListResponse(ctx *gin.Context, items interface{}, page int64, size int64) {
	ctx.JSON(http.StatusOK, ListResponseData{
		Code: http.StatusOK,
		Msg:  "success",
		Data: &Data{
			Items: items,
			Total: len(items.([]interface{})),
		},
		Page: page,
		Size: size,
	})
}

func (r *Response) ErrorResponseWithCode(ctx *gin.Context, code, httpCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
