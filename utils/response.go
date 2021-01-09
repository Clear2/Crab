package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR = -1
	SUCCESS = 0
)

type Response struct {
	Code int	`json:"code"`
	Data interface{} `json:"data"`
	Msg  string	`json:"msg"`
}

// ErrorItem 响应错误项
type ErrorItem struct {
	Code int `json:"code"`  // 错误码
	Message string	`json:"message"` // 错误信息
}

// ErrorResult 响应错误
type ErrorResult struct {
	Error ErrorItem	`json:"error"` // 错误项
}

func Result(c *gin.Context, code int, data interface{}, msg string)  {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})

}
