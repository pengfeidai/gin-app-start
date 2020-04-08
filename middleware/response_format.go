package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	Ctx *gin.Context
}

type response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Message interface{} `json:"msg"`
}

func (c *Context) Response(msg interface{}, content interface{}) {
	if msg == nil {
		c.Ctx.JSON(http.StatusOK, response{
			Code:    0,
			Success: true,
			Content: content,
			Message: msg,
		})
		return
	}
	// 错误格式
	c.Ctx.JSON(http.StatusOK, response{
		Code:    -1,
		Success: false,
		Content: content,
		Message: msg,
	})
}
