package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AppController struct {
	BaseController
	Ctx iris.Context
}

func (app *AppController) GetApps() mvc.Result {
	// token := app.Ctx.Values().Get(USERNAME).(*jwt.Token)
	// token := app.Ctx.GetHeader("Authorization").(*jwt.Token)
	token := app.Ctx.GetHeader("Authorization")

	return mvc.Response{Code: iris.StatusOK, Object: token}
}
