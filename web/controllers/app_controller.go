package controllers

import (
	"go_eden/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AppController struct {
	BaseController
	Ctx iris.Context
}

func (app *AppController) GetApps() mvc.Result {
	jwtInfo := app.Ctx.Values().Get("jwt").(*jwt.Token)
	golog.Debug(app.Ctx.JSON(jwtInfo))
	return mvc.Response{Code: iris.StatusOK, Object: model.Response{}}
}
