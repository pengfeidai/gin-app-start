package middleware

// ShouldBind
func (c *Context) Validate(p interface{}) error {
	if err := c.Ctx.ShouldBind(p); err != nil {
		logger.Error("param validate error:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindUri
func (c *Context) ValidateRouter(p interface{}) error {
	if err := c.Ctx.ShouldBindUri(p); err != nil {
		logger.Error("param validate err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindQuery
func (c *Context) ValidateQuery(p interface{}) error {
	if err := c.Ctx.ShouldBindQuery(p); err != nil {
		logger.Error("param validate err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}

// ShouldBindJSON
func (c *Context) ValidateJSON(p interface{}) error {
	if err := c.Ctx.ShouldBindJSON(p); err != nil {
		logger.Error("param validate err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}
