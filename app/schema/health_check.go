package schema

type Health struct {
	Name string `form:"name" validate:"required,min=1,max=32"`
}
