package schema

type Health struct {
	Name string `form:"name" binding:"required"`
}
