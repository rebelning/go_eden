package controllers

import (
	"go_eden/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ConfigController struct {
	BaseController
	Ctx     iris.Context
	Service service.ConfigService
}

func (c *ConfigController) GetMenulist() mvc.Result {

	resp := c.Service.GetMenuList()
	return mvc.Response{Code: iris.StatusOK, Object: resp}
}
