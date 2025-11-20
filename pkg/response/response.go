package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the standard API response structure
type Response struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id,omitempty" example:"trace-id-123"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// PageResponse represents a paginated response
type PageResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total" example:"100"`
	Page     int         `json:"page" example:"1"`
	PageSize int         `json:"page_size" example:"10"`
}

func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	pageResponse := PageResponse{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	Success(c, pageResponse)
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func ErrorWithTrace(c *gin.Context, code int, message string, traceID string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
		TraceID: traceID,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

