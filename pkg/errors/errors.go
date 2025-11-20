package errors

import "fmt"

type BusinessError struct {
	Code    int
	Message string
	Cause   error
}

func (e *BusinessError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("code: %d, message: %s, cause: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewBusinessError(code int, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

func WrapBusinessError(code int, message string, cause error) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

var (
	ErrInvalidParams = NewBusinessError(10001, "Invalid parameters")
	ErrUserNotFound  = NewBusinessError(10002, "User not found")
	ErrUnauthorized  = NewBusinessError(10003, "Unauthorized access")
	ErrUserExists    = NewBusinessError(10004, "User already exists")
	ErrDatabaseError = NewBusinessError(10005, "Database error")
	ErrInternalError = NewBusinessError(50000, "Internal server error")
)
