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

func (c *Context) Response(code int, msg interface{}, content interface{}) {
	if msg == nil {
		c.Ctx.JSON(http.StatusOK, response{
			Code:    code,
			Success: true,
			Content: content,
			Message: msg,
		})
		return
	}
	// 错误格式
	c.Ctx.JSON(http.StatusOK, response{
		Code:    code,
		Success: false,
		Content: content,
		Message: msg,
	})
}
