package controller

import (
	"gin-app-start/internal/dto"
	"gin-app-start/internal/service"
	"gin-app-start/pkg/errors"
	"gin-app-start/pkg/logger"
	"gin-app-start/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with username, email, phone and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateUserRequest	true	"User information"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 10001, "Parameter binding failed: "+err.Error())
		return
	}

	user, err := ctrl.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	response.Success(c, user)
}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get user information by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/api/v1/users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 10001, "Invalid user ID")
		return
	}

	user, err := ctrl.userService.GetUser(c.Request.Context(), uint(id))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	response.Success(c, user)
}

// UpdateUser godoc
//
//	@Summary		Update user information
//	@Description	Update user information by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			request	body		dto.UpdateUserRequest	true	"User information to update"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 10001, "Invalid user ID")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 10001, "Parameter binding failed: "+err.Error())
		return
	}

	user, err := ctrl.userService.UpdateUser(c.Request.Context(), uint(id), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	response.Success(c, user)
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete user by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/api/v1/users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 10001, "Invalid user ID")
		return
	}

	if err := ctrl.userService.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		handleServiceError(c, err)
		return
	}

	response.SuccessWithMessage(c, "Deleted successfully", nil)
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	Get paginated list of users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number"		default(1)
//	@Param			page_size	query		int	false	"Page size"			default(10)
//	@Success		200			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/api/v1/users [get]
func (ctrl *UserController) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := ctrl.userService.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	response.SuccessWithPage(c, users, total, page, pageSize)
}

func handleServiceError(c *gin.Context, err error) {
	var bizErr *errors.BusinessError
	if e, ok := err.(*errors.BusinessError); ok {
		bizErr = e
		response.Error(c, bizErr.Code, bizErr.Message)
	} else {
		logger.Error("Unknown error", zap.Error(err))
		response.Error(c, 50000, "Internal server error")
	}
}
