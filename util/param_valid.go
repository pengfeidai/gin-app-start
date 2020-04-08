package util

import (
	"log"
)

// var (
// 	Uni      *ut.UniversalTranslator
// 	Validate *validator.Validate
// )

// func InitValidate() {
// 	en := en.New()
// 	zh := zh.New()
// 	Uni = ut.New(en, zh)
// 	Validate = validator.New()
// }

func (c *Context) Validate(p interface{}) error {
	// validata := validator.New()
	// if err := validata.Struct(p); err != nil {
	// 	log.Println("param validate err:", err)
	// 	c.Response(err.Error(), nil)
	// 	return err
	// }
	if err := c.Ctx.ShouldBind(p); err != nil {
		log.Println("param validate err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}
