package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR = -1
	SUCEESS = 0
)

type Response struct {
	Code int 		 	`json:"code"`
	Data interface{} 	`json:"data"`
	Msg  string		 	`json:"msg"`
}


// ErrorItem 响应错误项
type ErrorItem struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}
// ErrorResult 响应错误
type ErrorResult struct {
	Error ErrorItem `json:"error"` // 错误项
}


func Result(c * gin.Context, code int, data interface{}, msg string)  {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
	fmt.Println(data)
}


// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	Result(c, SUCEESS, v, "操作成功")
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	Result(c, ERROR, err.Error(), "操作失败")
}
