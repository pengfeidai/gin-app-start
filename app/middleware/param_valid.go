package middleware

import (
	"gopkg.in/go-playground/validator.v9"
)

// ShouldBind
func (c *Context) Validate(p interface{}) error {
	// 参数绑定
	if err := c.Ctx.ShouldBind(p); err != nil {
		logger.Error("param ShouldBind error:", err)
		c.Response(400, err.Error(), nil)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		logger.Error("param validate error:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindUri
func (c *Context) ValidateRouter(p interface{}) error {
	if err := c.Ctx.ShouldBindUri(p); err != nil {
		logger.Error("param ShouldBindUri err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindQuery
func (c *Context) ValidateQuery(p interface{}) error {
	if err := c.Ctx.ShouldBindQuery(p); err != nil {
		logger.Error("param ShouldBindQuery err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		logger.Error("param validate error:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindJSON
func (c *Context) ValidateJSON(p interface{}) error {
	if err := c.Ctx.ShouldBindJSON(p); err != nil {
		logger.Error("param ShouldBindJSON err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		logger.Error("param validate error:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}
